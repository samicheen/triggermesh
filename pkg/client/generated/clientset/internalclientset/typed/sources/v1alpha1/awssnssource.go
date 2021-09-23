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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	scheme "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSSNSSourcesGetter has a method to return a AWSSNSSourceInterface.
// A group's client should implement this interface.
type AWSSNSSourcesGetter interface {
	AWSSNSSources(namespace string) AWSSNSSourceInterface
}

// AWSSNSSourceInterface has methods to work with AWSSNSSource resources.
type AWSSNSSourceInterface interface {
	Create(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.CreateOptions) (*v1alpha1.AWSSNSSource, error)
	Update(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.UpdateOptions) (*v1alpha1.AWSSNSSource, error)
	UpdateStatus(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.UpdateOptions) (*v1alpha1.AWSSNSSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AWSSNSSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AWSSNSSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSSNSSource, err error)
	AWSSNSSourceExpansion
}

// aWSSNSSources implements AWSSNSSourceInterface
type aWSSNSSources struct {
	client rest.Interface
	ns     string
}

// newAWSSNSSources returns a AWSSNSSources
func newAWSSNSSources(c *SourcesV1alpha1Client, namespace string) *aWSSNSSources {
	return &aWSSNSSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSSNSSource, and returns the corresponding aWSSNSSource object, and an error if there is any.
func (c *aWSSNSSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSSNSSource, err error) {
	result = &v1alpha1.AWSSNSSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssnssources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSSNSSources that match those selectors.
func (c *aWSSNSSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSSNSSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AWSSNSSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssnssources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSSNSSources.
func (c *aWSSNSSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssnssources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aWSSNSSource and creates it.  Returns the server's representation of the aWSSNSSource, and an error, if there is any.
func (c *aWSSNSSources) Create(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.CreateOptions) (result *v1alpha1.AWSSNSSource, err error) {
	result = &v1alpha1.AWSSNSSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssnssources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSSNSSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aWSSNSSource and updates it. Returns the server's representation of the aWSSNSSource, and an error, if there is any.
func (c *aWSSNSSources) Update(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.UpdateOptions) (result *v1alpha1.AWSSNSSource, err error) {
	result = &v1alpha1.AWSSNSSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssnssources").
		Name(aWSSNSSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSSNSSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aWSSNSSources) UpdateStatus(ctx context.Context, aWSSNSSource *v1alpha1.AWSSNSSource, opts v1.UpdateOptions) (result *v1alpha1.AWSSNSSource, err error) {
	result = &v1alpha1.AWSSNSSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssnssources").
		Name(aWSSNSSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSSNSSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aWSSNSSource and deletes it. Returns an error if one occurs.
func (c *aWSSNSSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssnssources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSSNSSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssnssources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aWSSNSSource.
func (c *aWSSNSSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSSNSSource, err error) {
	result = &v1alpha1.AWSSNSSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssnssources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}