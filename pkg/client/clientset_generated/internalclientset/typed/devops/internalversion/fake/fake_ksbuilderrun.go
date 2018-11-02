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

// FakeKsBuilderRuns implements KsBuilderRunInterface
type FakeKsBuilderRuns struct {
	Fake *FakeDevops
	ns   string
}

var ksbuilderrunsResource = schema.GroupVersionResource{Group: "devops.kubesphere.io", Version: "", Resource: "ksbuilderruns"}

var ksbuilderrunsKind = schema.GroupVersionKind{Group: "devops.kubesphere.io", Version: "", Kind: "KsBuilderRun"}

// Get takes name of the ksBuilderRun, and returns the corresponding ksBuilderRun object, and an error if there is any.
func (c *FakeKsBuilderRuns) Get(name string, options v1.GetOptions) (result *devops.KsBuilderRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ksbuilderrunsResource, c.ns, name), &devops.KsBuilderRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilderRun), err
}

// List takes label and field selectors, and returns the list of KsBuilderRuns that match those selectors.
func (c *FakeKsBuilderRuns) List(opts v1.ListOptions) (result *devops.KsBuilderRunList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ksbuilderrunsResource, ksbuilderrunsKind, c.ns, opts), &devops.KsBuilderRunList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &devops.KsBuilderRunList{}
	for _, item := range obj.(*devops.KsBuilderRunList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ksBuilderRuns.
func (c *FakeKsBuilderRuns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ksbuilderrunsResource, c.ns, opts))

}

// Create takes the representation of a ksBuilderRun and creates it.  Returns the server's representation of the ksBuilderRun, and an error, if there is any.
func (c *FakeKsBuilderRuns) Create(ksBuilderRun *devops.KsBuilderRun) (result *devops.KsBuilderRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ksbuilderrunsResource, c.ns, ksBuilderRun), &devops.KsBuilderRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilderRun), err
}

// Update takes the representation of a ksBuilderRun and updates it. Returns the server's representation of the ksBuilderRun, and an error, if there is any.
func (c *FakeKsBuilderRuns) Update(ksBuilderRun *devops.KsBuilderRun) (result *devops.KsBuilderRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ksbuilderrunsResource, c.ns, ksBuilderRun), &devops.KsBuilderRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilderRun), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKsBuilderRuns) UpdateStatus(ksBuilderRun *devops.KsBuilderRun) (*devops.KsBuilderRun, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ksbuilderrunsResource, "status", c.ns, ksBuilderRun), &devops.KsBuilderRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilderRun), err
}

// Delete takes name of the ksBuilderRun and deletes it. Returns an error if one occurs.
func (c *FakeKsBuilderRuns) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ksbuilderrunsResource, c.ns, name), &devops.KsBuilderRun{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKsBuilderRuns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ksbuilderrunsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &devops.KsBuilderRunList{})
	return err
}

// Patch applies the patch and returns the patched ksBuilderRun.
func (c *FakeKsBuilderRuns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *devops.KsBuilderRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ksbuilderrunsResource, c.ns, name, data, subresources...), &devops.KsBuilderRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*devops.KsBuilderRun), err
}
