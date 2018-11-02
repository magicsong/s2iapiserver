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
*/package v1alpha1

import (
	v1alpha1 "github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	scheme "github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KsBuildersGetter has a method to return a KsBuilderInterface.
// A group's client should implement this interface.
type KsBuildersGetter interface {
	KsBuilders(namespace string) KsBuilderInterface
}

// KsBuilderInterface has methods to work with KsBuilder resources.
type KsBuilderInterface interface {
	Create(*v1alpha1.KsBuilder) (*v1alpha1.KsBuilder, error)
	Update(*v1alpha1.KsBuilder) (*v1alpha1.KsBuilder, error)
	UpdateStatus(*v1alpha1.KsBuilder) (*v1alpha1.KsBuilder, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KsBuilder, error)
	List(opts v1.ListOptions) (*v1alpha1.KsBuilderList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KsBuilder, err error)
	KsBuilderExpansion
}

// ksBuilders implements KsBuilderInterface
type ksBuilders struct {
	client rest.Interface
	ns     string
}

// newKsBuilders returns a KsBuilders
func newKsBuilders(c *DevopsV1alpha1Client, namespace string) *ksBuilders {
	return &ksBuilders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ksBuilder, and returns the corresponding ksBuilder object, and an error if there is any.
func (c *ksBuilders) Get(name string, options v1.GetOptions) (result *v1alpha1.KsBuilder, err error) {
	result = &v1alpha1.KsBuilder{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ksbuilders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KsBuilders that match those selectors.
func (c *ksBuilders) List(opts v1.ListOptions) (result *v1alpha1.KsBuilderList, err error) {
	result = &v1alpha1.KsBuilderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ksbuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ksBuilders.
func (c *ksBuilders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ksbuilders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a ksBuilder and creates it.  Returns the server's representation of the ksBuilder, and an error, if there is any.
func (c *ksBuilders) Create(ksBuilder *v1alpha1.KsBuilder) (result *v1alpha1.KsBuilder, err error) {
	result = &v1alpha1.KsBuilder{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ksbuilders").
		Body(ksBuilder).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ksBuilder and updates it. Returns the server's representation of the ksBuilder, and an error, if there is any.
func (c *ksBuilders) Update(ksBuilder *v1alpha1.KsBuilder) (result *v1alpha1.KsBuilder, err error) {
	result = &v1alpha1.KsBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ksbuilders").
		Name(ksBuilder.Name).
		Body(ksBuilder).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ksBuilders) UpdateStatus(ksBuilder *v1alpha1.KsBuilder) (result *v1alpha1.KsBuilder, err error) {
	result = &v1alpha1.KsBuilder{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ksbuilders").
		Name(ksBuilder.Name).
		SubResource("status").
		Body(ksBuilder).
		Do().
		Into(result)
	return
}

// Delete takes name of the ksBuilder and deletes it. Returns an error if one occurs.
func (c *ksBuilders) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ksbuilders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ksBuilders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ksbuilders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ksBuilder.
func (c *ksBuilders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KsBuilder, err error) {
	result = &v1alpha1.KsBuilder{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ksbuilders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
