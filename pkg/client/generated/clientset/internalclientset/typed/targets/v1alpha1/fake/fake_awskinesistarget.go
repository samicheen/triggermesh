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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAWSKinesisTargets implements AWSKinesisTargetInterface
type FakeAWSKinesisTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var awskinesistargetsResource = schema.GroupVersionResource{Group: "targets.triggermesh.io", Version: "v1alpha1", Resource: "awskinesistargets"}

var awskinesistargetsKind = schema.GroupVersionKind{Group: "targets.triggermesh.io", Version: "v1alpha1", Kind: "AWSKinesisTarget"}

// Get takes name of the aWSKinesisTarget, and returns the corresponding aWSKinesisTarget object, and an error if there is any.
func (c *FakeAWSKinesisTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSKinesisTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awskinesistargetsResource, c.ns, name), &v1alpha1.AWSKinesisTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSKinesisTarget), err
}

// List takes label and field selectors, and returns the list of AWSKinesisTargets that match those selectors.
func (c *FakeAWSKinesisTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSKinesisTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awskinesistargetsResource, awskinesistargetsKind, c.ns, opts), &v1alpha1.AWSKinesisTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSKinesisTargetList{ListMeta: obj.(*v1alpha1.AWSKinesisTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSKinesisTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSKinesisTargets.
func (c *FakeAWSKinesisTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awskinesistargetsResource, c.ns, opts))

}

// Create takes the representation of a aWSKinesisTarget and creates it.  Returns the server's representation of the aWSKinesisTarget, and an error, if there is any.
func (c *FakeAWSKinesisTargets) Create(ctx context.Context, aWSKinesisTarget *v1alpha1.AWSKinesisTarget, opts v1.CreateOptions) (result *v1alpha1.AWSKinesisTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awskinesistargetsResource, c.ns, aWSKinesisTarget), &v1alpha1.AWSKinesisTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSKinesisTarget), err
}

// Update takes the representation of a aWSKinesisTarget and updates it. Returns the server's representation of the aWSKinesisTarget, and an error, if there is any.
func (c *FakeAWSKinesisTargets) Update(ctx context.Context, aWSKinesisTarget *v1alpha1.AWSKinesisTarget, opts v1.UpdateOptions) (result *v1alpha1.AWSKinesisTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awskinesistargetsResource, c.ns, aWSKinesisTarget), &v1alpha1.AWSKinesisTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSKinesisTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSKinesisTargets) UpdateStatus(ctx context.Context, aWSKinesisTarget *v1alpha1.AWSKinesisTarget, opts v1.UpdateOptions) (*v1alpha1.AWSKinesisTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awskinesistargetsResource, "status", c.ns, aWSKinesisTarget), &v1alpha1.AWSKinesisTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSKinesisTarget), err
}

// Delete takes name of the aWSKinesisTarget and deletes it. Returns an error if one occurs.
func (c *FakeAWSKinesisTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awskinesistargetsResource, c.ns, name), &v1alpha1.AWSKinesisTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSKinesisTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awskinesistargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSKinesisTargetList{})
	return err
}

// Patch applies the patch and returns the patched aWSKinesisTarget.
func (c *FakeAWSKinesisTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSKinesisTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awskinesistargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSKinesisTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSKinesisTarget), err
}