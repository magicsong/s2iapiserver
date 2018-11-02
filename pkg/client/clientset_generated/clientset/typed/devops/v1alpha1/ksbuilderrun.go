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

// KsBuilderRunsGetter has a method to return a KsBuilderRunInterface.
// A group's client should implement this interface.
type KsBuilderRunsGetter interface {
	KsBuilderRuns(namespace string) KsBuilderRunInterface
}

// KsBuilderRunInterface has methods to work with KsBuilderRun resources.
type KsBuilderRunInterface interface {
	Create(*v1alpha1.KsBuilderRun) (*v1alpha1.KsBuilderRun, error)
	Update(*v1alpha1.KsBuilderRun) (*v1alpha1.KsBuilderRun, error)
	UpdateStatus(*v1alpha1.KsBuilderRun) (*v1alpha1.KsBuilderRun, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KsBuilderRun, error)
	List(opts v1.ListOptions) (*v1alpha1.KsBuilderRunList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KsBuilderRun, err error)
	KsBuilderRunExpansion
}

// ksBuilderRuns implements KsBuilderRunInterface
type ksBuilderRuns struct {
	client rest.Interface
	ns     string
}

// newKsBuilderRuns returns a KsBuilderRuns
func newKsBuilderRuns(c *DevopsV1alpha1Client, namespace string) *ksBuilderRuns {
	return &ksBuilderRuns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ksBuilderRun, and returns the corresponding ksBuilderRun object, and an error if there is any.
func (c *ksBuilderRuns) Get(name string, options v1.GetOptions) (result *v1alpha1.KsBuilderRun, err error) {
	result = &v1alpha1.KsBuilderRun{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KsBuilderRuns that match those selectors.
func (c *ksBuilderRuns) List(opts v1.ListOptions) (result *v1alpha1.KsBuilderRunList, err error) {
	result = &v1alpha1.KsBuilderRunList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ksBuilderRuns.
func (c *ksBuilderRuns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a ksBuilderRun and creates it.  Returns the server's representation of the ksBuilderRun, and an error, if there is any.
func (c *ksBuilderRuns) Create(ksBuilderRun *v1alpha1.KsBuilderRun) (result *v1alpha1.KsBuilderRun, err error) {
	result = &v1alpha1.KsBuilderRun{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		Body(ksBuilderRun).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ksBuilderRun and updates it. Returns the server's representation of the ksBuilderRun, and an error, if there is any.
func (c *ksBuilderRuns) Update(ksBuilderRun *v1alpha1.KsBuilderRun) (result *v1alpha1.KsBuilderRun, err error) {
	result = &v1alpha1.KsBuilderRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		Name(ksBuilderRun.Name).
		Body(ksBuilderRun).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ksBuilderRuns) UpdateStatus(ksBuilderRun *v1alpha1.KsBuilderRun) (result *v1alpha1.KsBuilderRun, err error) {
	result = &v1alpha1.KsBuilderRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		Name(ksBuilderRun.Name).
		SubResource("status").
		Body(ksBuilderRun).
		Do().
		Into(result)
	return
}

// Delete takes name of the ksBuilderRun and deletes it. Returns an error if one occurs.
func (c *ksBuilderRuns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ksBuilderRuns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ksbuilderruns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ksBuilderRun.
func (c *ksBuilderRuns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KsBuilderRun, err error) {
	result = &v1alpha1.KsBuilderRun{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ksbuilderruns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
