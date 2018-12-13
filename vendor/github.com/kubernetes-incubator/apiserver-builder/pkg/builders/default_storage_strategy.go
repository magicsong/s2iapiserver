/*
Copyright 2017 The Kubernetes Authors.

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

package builders

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

var _ rest.RESTCreateStrategy = &DefaultStorageStrategy{}
var _ rest.RESTDeleteStrategy = &DefaultStorageStrategy{}
var _ rest.RESTUpdateStrategy = &DefaultStorageStrategy{}

var StorageStrategySingleton = DefaultStorageStrategy{
	Scheme,
	names.SimpleNameGenerator,
}

type DefaultStorageStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (DefaultStorageStrategy) ObjectNameFunc(obj runtime.Object) (string, error) {
	switch obj := obj.(type) {
	default:
		return "", fmt.Errorf(
			"Cannot get name for object type %T.  Must implement HasObjectMeta or define "+
				"its own ObjectNameFunc in its storage strategy.", obj)
	case HasObjectMeta:
		// Get the name from the metadata
		return obj.GetObjectMeta().Name, nil
	}
}

// Build sets the strategy for the store
func (DefaultStorageStrategy) Build(builder StorageBuilder, store *StorageWrapper, options *generic.StoreOptions) {
	store.PredicateFunc = builder.BasicMatch
	store.ObjectNameFunc = builder.ObjectNameFunc
	store.CreateStrategy = builder
	store.UpdateStrategy = builder
	store.DeleteStrategy = builder
	store.TableConvertor = builder
	store.ShortNamesProvider = builder

	options.AttrFunc = builder.GetAttrs
	options.TriggerFunc = builder.TriggerFunc
}

func (DefaultStorageStrategy) NamespaceScoped() bool { return true }

func (DefaultStorageStrategy) AllowCreateOnUpdate() bool { return true }

func (DefaultStorageStrategy) AllowUnconditionalUpdate() bool { return true }

func (DefaultStorageStrategy) Canonicalize(obj runtime.Object) {}

func (DefaultStorageStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	switch t := obj.(type) {
	default:
	case HasObjectMetaSpecStatus:
		// Clear the status if the resource has a Status
		t.GetObjectMeta().Generation = 1
		t.SetStatus(t.NewStatus())
	}
}

func (DefaultStorageStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	// Don't update the status if the resource has a Status
	switch n := obj.(type) {
	default:
	case HasObjectMetaSpecStatus:
		o := old.(HasObjectMetaSpecStatus)
		n.SetStatus(o.GetStatus())

		// Spec and annotation updates bump the generation.
		if !reflect.DeepEqual(n.GetSpec(), o.GetSpec()) ||
			!reflect.DeepEqual(n.GetObjectMeta().Annotations, o.GetObjectMeta().Annotations) {
			n.GetObjectMeta().Generation = o.GetObjectMeta().Generation + 1
		}
	}
}

func (DefaultStorageStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (DefaultStorageStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (b DefaultStorageStrategy) GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	switch t := obj.(type) {
	case HasObjectMeta:
		apiserver := obj.(HasObjectMeta)
		return labels.Set(apiserver.GetObjectMeta().Labels),
			b.GetSelectableFields(apiserver),
			apiserver.GetObjectMeta().Initializers != nil,
			nil
	default:
		return nil, nil, false, fmt.Errorf(
			"Cannot get attributes for object type %v which does not implement HasObjectMeta.", t)
	}
}

func (b DefaultStorageStrategy) TriggerFunc(obj runtime.Object) []storage.MatchValue {
	return []storage.MatchValue{}
}

// GetSelectableFields returns a field set that represents the object.
func (DefaultStorageStrategy) GetSelectableFields(obj HasObjectMeta) fields.Set {
	return generic.ObjectMetaFieldsSet(obj.GetObjectMeta(), true)
}

// MatchResource is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func (b DefaultStorageStrategy) BasicMatch(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: b.GetAttrs,
	}
}
func (DefaultStorageStrategy) ShortNames() []string {
	return nil
}

func (DefaultStorageStrategy) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1beta1.Table, error) {
	var table metav1beta1.Table
	var swaggerMetadataDescriptions = metav1.ObjectMeta{}.SwaggerDoc()
	fn := func(obj runtime.Object) error {
		m, err := meta.Accessor(obj)
		if err != nil {
			return fmt.Errorf("Not support TableConvertor interface")
		}
		table.Rows = append(table.Rows, metav1beta1.TableRow{
			Cells:  []interface{}{m.GetName(), m.GetCreationTimestamp().Time.UTC().Format(time.RFC3339)},
			Object: runtime.RawExtension{Object: obj},
		})
		return nil
	}
	switch {
	case meta.IsListType(object):
		if err := meta.EachListItem(object, fn); err != nil {
			return nil, err
		}
	default:
		if err := fn(object); err != nil {
			return nil, err
		}
	}
	if m, err := meta.ListAccessor(object); err == nil {
		table.ResourceVersion = m.GetResourceVersion()
		table.SelfLink = m.GetSelfLink()
		table.Continue = m.GetContinue()
	} else {
		if m, err := meta.CommonAccessor(object); err == nil {
			table.ResourceVersion = m.GetResourceVersion()
			table.SelfLink = m.GetSelfLink()
		}
	}
	table.ColumnDefinitions = []metav1beta1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: swaggerMetadataDescriptions["name"]},
		{Name: "Created At", Type: "date", Description: swaggerMetadataDescriptions["creationTimestamp"]},
	}
	return &table, nil
}

//
// Status Strategies
//

var StatusStorageStrategySingleton = DefaultStatusStorageStrategy{StorageStrategySingleton}

type DefaultStatusStorageStrategy struct {
	DefaultStorageStrategy
}

func (DefaultStatusStorageStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	switch n := obj.(type) {
	default:
	case HasObjectMetaSpecStatus:
		// Only update the Status
		o := old.(HasObjectMetaSpecStatus)
		n.SetSpec(o.GetSpec())
		n.GetObjectMeta().Labels = o.GetObjectMeta().Labels
	}
}
