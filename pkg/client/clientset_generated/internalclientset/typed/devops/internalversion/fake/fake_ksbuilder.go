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
*/package fake

import (
	devops "github.com/magicsong/s2iapiserver/pkg/apis/devops"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKsBuilders implements KsBuilderInterface
type FakeKsBuilders struct {
	Fake *FakeDevops
	ns   string
}

var ksbuildersResource = schema.GroupVersionResource{Group: "devops.kubesphere.io", Version: "", Resource: "ksbuilders"}

var ksbuildersKind = schema.GroupVersionKind{Group: "devops.kubesphere.io", Version: "", Kind: "KsBuilder"}

// Get takes name of the ksBuilder, and returns the corresponding ksBuilder object, and an error if there is any.
func (c *FakeKsBuilders) Get(name string, options v1.GetOptions) (result *devops.KsBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ksbuildersResource, c.ns, name), &devops.KsBuilder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilder), err
}

// List takes label and field selectors, and returns the list of KsBuilders that match those selectors.
func (c *FakeKsBuilders) List(opts v1.ListOptions) (result *devops.KsBuilderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ksbuildersResource, ksbuildersKind, c.ns, opts), &devops.KsBuilderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &devops.KsBuilderList{}
	for _, item := range obj.(*devops.KsBuilderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ksBuilders.
func (c *FakeKsBuilders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ksbuildersResource, c.ns, opts))

}

// Create takes the representation of a ksBuilder and creates it.  Returns the server's representation of the ksBuilder, and an error, if there is any.
func (c *FakeKsBuilders) Create(ksBuilder *devops.KsBuilder) (result *devops.KsBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ksbuildersResource, c.ns, ksBuilder), &devops.KsBuilder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilder), err
}

// Update takes the representation of a ksBuilder and updates it. Returns the server's representation of the ksBuilder, and an error, if there is any.
func (c *FakeKsBuilders) Update(ksBuilder *devops.KsBuilder) (result *devops.KsBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ksbuildersResource, c.ns, ksBuilder), &devops.KsBuilder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKsBuilders) UpdateStatus(ksBuilder *devops.KsBuilder) (*devops.KsBuilder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ksbuildersResource, "status", c.ns, ksBuilder), &devops.KsBuilder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilder), err
}

// Delete takes name of the ksBuilder and deletes it. Returns an error if one occurs.
func (c *FakeKsBuilders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ksbuildersResource, c.ns, name), &devops.KsBuilder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKsBuilders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ksbuildersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &devops.KsBuilderList{})
	return err
}

// Patch applies the patch and returns the patched ksBuilder.
func (c *FakeKsBuilders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *devops.KsBuilder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ksbuildersResource, c.ns, name, data, subresources...), &devops.KsBuilder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilder), err
}
