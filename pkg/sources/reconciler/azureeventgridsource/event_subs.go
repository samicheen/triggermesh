/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azureeventgridsource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/pkg/controller"
	"knative.dev/pkg/reconciler"

	azureeventgrid "github.com/Azure/azure-sdk-for-go/profiles/latest/eventgrid/mgmt/eventgrid"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/sources/client/azure/eventgrid"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common/event"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common/skip"
)

const crudTimeout = time.Second * 15

const (
	defaultMaxDeliveryAttempts = 30
	defaultEventTTL            = 1440
)

// ensureEventSubscription ensures an event subscription exists with the expected configuration.
// Required permissions:
//  - Microsoft.EventGrid/eventSubscriptions/read
//  - Microsoft.EventGrid/eventSubscriptions/write
//  - Microsoft.EventHub/namespaces/eventhubs/write
func ensureEventSubscription(ctx context.Context, cli eventgrid.EventSubscriptionsClient, eventHubResID string) error {
	if skip.Skip(ctx) {
		return nil
	}

	src := v1alpha1.SourceFromContext(ctx)
	typedSrc := src.(*v1alpha1.AzureEventGridSource)

	status := &typedSrc.Status

	// read current event subscription

	scope := typedSrc.Spec.Scope.String()
	subsName := subscriptionName(typedSrc)

	restCtx, cancel := context.WithTimeout(ctx, crudTimeout)
	defer cancel()

	currentEventSubs, err := cli.Get(restCtx, scope, subsName)
	switch {
	case isNotFound(err):
		// no-op
	case isDenied(err):
		status.MarkNotSubscribed(v1alpha1.AzureReasonAPIError, "Access denied to event subscription API: "+toErrMsg(err))
		return controller.NewPermanentError(failGetEventSubscriptionEvent(scope, err))
	case err != nil:
		status.MarkNotSubscribed(v1alpha1.AzureReasonAPIError, "Cannot look up event subscription: "+toErrMsg(err))
		return fmt.Errorf("%w", failGetEventSubscriptionEvent(scope, err))
	}

	subsExists := currentEventSubs.ID != nil

	// compare and create/update event subscription

	desiredEventSubs := newEventSubscription(eventHubResID, typedSrc.GetEventTypes())

	if equalEventSubscription(ctx, desiredEventSubs, currentEventSubs) {
		eventSubscriptionResID, err := parseEventSubscriptionResID(*currentEventSubs.ID)
		if err != nil {
			return fmt.Errorf("converting resource ID string to structured resource ID: %w", err)
		}

		status.EventSubscriptionID = eventSubscriptionResID
		status.MarkSubscribed()
		return nil
	}

	restCtx, cancel = context.WithTimeout(ctx, crudTimeout)
	defer cancel()

	_, err = cli.CreateOrUpdate(restCtx, scope, subsName, desiredEventSubs)
	switch {
	case isDenied(err):
		status.MarkNotSubscribed(v1alpha1.AzureReasonAPIError, "Access denied to event subscription API: "+toErrMsg(err))
		return controller.NewPermanentError(failSubscribeEvent(scope, subsExists, err))
	case err != nil:
		status.MarkNotSubscribed(v1alpha1.AzureReasonAPIError, "Cannot subscribe to events: "+toErrMsg(err))
		return fmt.Errorf("%w", failSubscribeEvent(scope, subsExists, err))
	}

	recordSubscribedEvent(ctx, subsName, scope, subsExists)

	// NOTE(antoineco): CreateOrUpdate() returns a "future" instead of the
	// actual Event Grid subscription, so setting status.EventSubscriptionID
	// here would require waiting for the async operation (create/update)
	// to complete, which might take several seconds.
	// Because reporting the subscription ID quickly isn't essential, we
	// prefer to return early and accept that the subscription ID will only
	// be propagated in the status during the next reconciliation.

	status.MarkSubscribed()

	return nil
}

// newEventSubscription returns the desired state of the event subscription for
// the given source.
func newEventSubscription(eventHubResID string, eventTypes []string) azureeventgrid.EventSubscription {
	// Fields marked with a '*' below are attributes which would be
	// defaulted on creation by Azure if not explicitly set, but which we
	// set manually nevertheless in order to ease the comparison with the
	// current state in the main synchronization logic.

	return azureeventgrid.EventSubscription{
		EventSubscriptionProperties: &azureeventgrid.EventSubscriptionProperties{
			Destination: azureeventgrid.EventHubEventSubscriptionDestination{
				EndpointType: azureeventgrid.EndpointTypeEventHub,
				EventHubEventSubscriptionDestinationProperties: &azureeventgrid.EventHubEventSubscriptionDestinationProperties{
					ResourceID: to.StringPtr(eventHubResID),
				},
			},
			Filter: &azureeventgrid.EventSubscriptionFilter{
				IncludedEventTypes: to.StringSlicePtr(eventTypes),
				SubjectBeginsWith:  to.StringPtr(""), // *
				SubjectEndsWith:    to.StringPtr(""), // *
			},
			RetryPolicy: &azureeventgrid.RetryPolicy{
				MaxDeliveryAttempts:      to.Int32Ptr(defaultMaxDeliveryAttempts), // *
				EventTimeToLiveInMinutes: to.Int32Ptr(defaultEventTTL),            // *
			},
			EventDeliverySchema: azureeventgrid.CloudEventSchemaV10,
		},
	}
}

// ensureNoEventSubscription ensures the event subscription is removed.
// Required permissions:
//  - Microsoft.EventGrid/eventSubscriptions/delete
func ensureNoEventSubscription(ctx context.Context, cli eventgrid.EventSubscriptionsClient) reconciler.Event {
	if skip.Skip(ctx) {
		return nil
	}

	src := v1alpha1.SourceFromContext(ctx)
	typedSrc := src.(*v1alpha1.AzureEventGridSource)

	scope := typedSrc.Spec.Scope.String()
	subsName := subscriptionName(typedSrc)

	restCtx, cancel := context.WithTimeout(ctx, crudTimeout)
	defer cancel()

	_, err := cli.Delete(restCtx, scope, subsName)
	switch {
	case isNotFound(err):
		event.Warn(ctx, ReasonUnsubscribed, "Event subscription not found, skipping deletion")
		return nil
	case isDenied(err):
		// it is unlikely that we recover from auth errors in the
		// finalizer, so we simply record a warning event and return
		event.Warn(ctx, ReasonFailedUnsubscribe,
			"Access denied to event subscription API. Ignoring: %s", toErrMsg(err))
		return nil
	case err != nil:
		return failUnsubscribeEvent(scope, err)
	}

	event.Normal(ctx, ReasonUnsubscribed, "Deleted event subscription %q for Azure resource %q",
		subsName, scope)

	return nil
}

// parseEventSubscriptionResID parses the given Event Hub resource ID string to
// a structured resource ID.
func parseEventSubscriptionResID(resIDStr string) (*v1alpha1.AzureResourceID, error) {
	resID := &v1alpha1.AzureResourceID{}

	err := json.Unmarshal([]byte(strconv.Quote(resIDStr)), resID)
	if err != nil {
		return nil, fmt.Errorf("deserializing resource ID string: %w", err)
	}

	return resID, nil
}

// toErrMsg returns the given error as a string.
// If the error is an Azure API error, the error message is sanitized while
// still preserving the concatenation of all nested levels of errors.
//
// Used to remove clutter from errors before writing them to status conditions.
func toErrMsg(err error) string {
	return recursErrMsg("", err)
}

// recursErrMsg concatenates the messages of deeply nested API errors recursively.
func recursErrMsg(errMsg string, err error) string {
	if errMsg != "" {
		errMsg += ": "
	}

	switch tErr := err.(type) {
	case autorest.DetailedError:
		return recursErrMsg(errMsg+tErr.Message, tErr.Original)
	case *azure.RequestError:
		if tErr.DetailedError.Original != nil {
			return recursErrMsg(errMsg+tErr.DetailedError.Message, tErr.DetailedError.Original)
		}
		if tErr.ServiceError != nil {
			return errMsg + tErr.ServiceError.Message
		}
	case adal.TokenRefreshError:
		// This type of error is returned when the OAuth authentication with Azure Active Directory fails, often
		// due to an invalid or expired secret.
		//
		// The associated message is typically opaque and contains elements that are unique to each request
		// (trace/correlation IDs, timestamps), which causes an infinite loop of reconciliation if propagated to
		// the object's status conditions.
		// Instead of resorting to over-engineered error parsing techniques to get around the verbosity of the
		// message, we simply return a short and generic error description.
		return errMsg + "Invalid client secret"
	}

	return errMsg + err.Error()
}

// isNotFound returns whether the given error indicates that some Azure
// resource was not found.
func isNotFound(err error) bool {
	if dErr := (autorest.DetailedError{}); errors.As(err, &dErr) {
		return dErr.StatusCode == http.StatusNotFound
	}
	return false
}

// isDenied returns whether the given error indicates that a request to the
// Azure API could not be authorized.
// This category of issues is unrecoverable without user intervention.
func isDenied(err error) bool {
	if dErr := (autorest.DetailedError{}); errors.As(err, &dErr) {
		if code, ok := dErr.StatusCode.(int); ok {
			return code == http.StatusUnauthorized || code == http.StatusForbidden
		}
	}

	return false
}

// subscriptionName returns a predictable name for an Event Grid event
// subscription associated with the given source instance.
func subscriptionName(o *v1alpha1.AzureEventGridSource) string {
	return "io.triggermesh.azureeventgridsources." + o.Namespace + "." + o.Name
}

// recordSubscribedEvent records a Kubernetes API event which indicates that an
// event subscription was either created or updated.
func recordSubscribedEvent(ctx context.Context, subsName, resource string, isUpdate bool) {
	verb := "Created"
	if isUpdate {
		verb = "Updated"
	}

	event.Normal(ctx, ReasonSubscribed, "%s event subscription %q for Azure resource %q",
		verb, subsName, resource)
}

// failGetEventSubscriptionEvent returns a reconciler event which indicates
// that an event subscription for the given Azure resource could not be
// retrieved from the Azure API.
func failGetEventSubscriptionEvent(resource string, origErr error) reconciler.Event {
	return reconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedSubscribe,
		"Error getting event subscription for Azure resource %q: %s", resource, toErrMsg(origErr))
}

// failSubscribeEvent returns a reconciler event which indicates that an event
// subscription for the given Azure resource could not be created or updated
// via the Azure API.
func failSubscribeEvent(resource string, isUpdate bool, origErr error) reconciler.Event {
	verb := "creating"
	if isUpdate {
		verb = "updating"
	}

	return reconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedSubscribe,
		"Error %s event subscription for Azure resource %q: %s", verb, resource, toErrMsg(origErr))
}

// failUnsubscribeEvent returns a reconciler event which indicates that an
// event subscription for the given Azure resource could not be deleted via
// the Azure API.
func failUnsubscribeEvent(resource string, origErr error) reconciler.Event {
	return reconciler.NewEvent(corev1.EventTypeWarning, ReasonFailedSubscribe,
		"Error deleting event subscription for Azure resource %q: %s", resource, toErrMsg(origErr))
}