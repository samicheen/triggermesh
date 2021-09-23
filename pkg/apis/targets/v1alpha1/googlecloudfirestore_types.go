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

package v1alpha1

import (
	"github.com/triggermesh/triggermesh/pkg/apis/targets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleCloudFirestoreTarget is the Schema for the GoogleCloudFirestore Target.
type GoogleCloudFirestoreTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GoogleCloudFirestoreTargetSpec   `json:"spec"`
	Status GoogleCloudFirestoreTargetStatus `json:"status,omitempty"`
}

// Check the interfaces GoogleCloudFirestoreTarget should be implementing.
var (
	_ runtime.Object            = (*GoogleCloudFirestoreTarget)(nil)
	_ kmeta.OwnerRefable        = (*GoogleCloudFirestoreTarget)(nil)
	_ targets.IntegrationTarget = (*GoogleCloudFirestoreTarget)(nil)
	_ targets.EventSource       = (*GoogleCloudFirestoreTarget)(nil)
	_ duckv1.KRShaped           = (*GoogleCloudFirestoreTarget)(nil)
)

// GoogleCloudFirestoreTargetSpec holds the desired state of the GoogleCloudFirestoreTarget.
type GoogleCloudFirestoreTargetSpec struct {

	// Credentials represents how Google Firestore credentials should be provided in the secret
	Credentials SecretValueFromSource `json:"credentialsJson"`

	// DefaultCollection sets a default Firestore collection to select from
	DefaultCollection string `json:"defaultCollection"`

	// ProjectID specifies the Google project ID
	ProjectID string `json:"projectID"`

	// Whether to omit CloudEvent context attributes in created SNS notifications.
	// When this property is false (default), the entire CloudEvent payload is included.
	// When this property is true, only the CloudEvent data is included.
	DiscardCEContext bool `json:"discardCloudEventContext"`

	// EventOptions for targets
	EventOptions *EventOptions `json:"eventOptions,omitempty"`
}

// GoogleCloudFirestoreTargetStatus communicates the observed state of the GoogleCloudFirestoreTarget (from the controller).
type GoogleCloudFirestoreTargetStatus struct {
	duckv1.Status        `json:",inline"`
	duckv1.AddressStatus `json:",inline"`
	CloudEventStatus     `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleCloudFirestoreTargetList is a list of GoogleCloudFirestoreTarget resources
type GoogleCloudFirestoreTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GoogleCloudFirestoreTarget `json:"items"`
}