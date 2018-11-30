/*
Copyright 2018 The Kubesphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/package internalversion

import (
	devops "github.com/magicsong/s2iapiserver/pkg/apis/devops"
	scheme "github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// S2iBuildersGetter has a method to return a S2iBuilderInterface.
// A group's client should implement this interface.
type S2iBuildersGetter interface {
	S2iBuilders(namespace string) S2iBuilderInterface
}

// S2iBuilderInterface has methods to work with S2iBuilder resources.
type S2iBuilderInterface interface {
	Create(*devops.S2iBuilder) (*devops.S2iBuilder, error)
	Update(*devops.S2iBuilder) (*devops.S2iBuilder, error)
	UpdateStatus(*devops.S2iBuilder) (*devops.S2iBuilder, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*devops.S2iBuilder, error)
	List(opts v1.ListOptions) (*devops.S2iBuilderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *devops.S2iBuilder, err error)
	S2iBuilderExpansion
}

// s2iBuilders implements S2iBuilderInterface
type s2iBuilders struct {
	client rest.Interface
	ns     string
}

// newS2iBuilders returns a S2iBuilders
func newS2iBuilders(c *DevopsClient, namespace string) *s2iBuilders {
	return &s2iBuilders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the s2iBuilder, and returns the corresponding s2iBuilder object, and an error if there is any.
func (c *s2iBuilders) Get(name string, options v1.GetOptions) (result *devops.S2iBuilder, err error) {
	result = &devops.S2iBuilder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2ibuilders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of S2iBuilders that match those selectors.
func (c *s2iBuilders) List(opts v1.ListOptions) (result *devops.S2iBuilderList, err error) {
	result = &devops.S2iBuilderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("s2ibuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested s2iBuilders.
func (c *s2iBuilders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("s2ibuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a s2iBuilder and creates it.  Returns the server's representation of the s2iBuilder, and an error, if there is any.
func (c *s2iBuilders) Create(s2iBuilder *devops.S2iBuilder) (result *devops.S2iBuilder, err error) {
	result = &devops.S2iBuilder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("s2ibuilders").
		Body(s2iBuilder).
		Do().
		Into(result)
	return
}

// Update takes the representation of a s2iBuilder and updates it. Returns the server's representation of the s2iBuilder, and an error, if there is any.
func (c *s2iBuilders) Update(s2iBuilder *devops.S2iBuilder) (result *devops.S2iBuilder, err error) {
	result = &devops.S2iBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2ibuilders").
		Name(s2iBuilder.Name).
		Body(s2iBuilder).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *s2iBuilders) UpdateStatus(s2iBuilder *devops.S2iBuilder) (result *devops.S2iBuilder, err error) {
	result = &devops.S2iBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("s2ibuilders").
		Name(s2iBuilder.Name).
		SubResource("status").
		Body(s2iBuilder).
		Do().
		Into(result)
	return
}

// Delete takes name of the s2iBuilder and deletes it. Returns an error if one occurs.
func (c *s2iBuilders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2ibuilders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *s2iBuilders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("s2ibuilders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched s2iBuilder.
func (c *s2iBuilders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *devops.S2iBuilder, err error) {
	result = &devops.S2iBuilder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("s2ibuilders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
