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
*/

package v1alpha1

import (
	"context"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type Rerun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

var _ rest.CreaterUpdater = &RerunS2iRunREST{}
var _ rest.Patcher = &RerunS2iRunREST{}

// +k8s:deepcopy-gen=false
type RerunS2iRunREST struct {
	Registry devops.S2iRunRegistry
}

const (
	RerunAnnotationKey = "v1alpha1.devops.kubesphere.io/rerun"
	RerunValue         = "yes"
)

func (r *RerunS2iRunREST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	sub := obj.(*Rerun)
	rec, err := r.Registry.GetS2iRun(ctx, sub.Name, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	// Modify rec in someway before writing it back to storage
	if rec.Annotations == nil {
		rec.Annotations = make(map[string]string)
		rec.Annotations[RerunAnnotationKey] = RerunValue
	} else {
		rec.Annotations[RerunAnnotationKey] = RerunValue
	}
	r.Registry.UpdateS2iRun(ctx, rec)
	return rec, nil
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *RerunS2iRunREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return &Rerun{}, nil
}

// Update alters the status subset of an object.
func (r *RerunS2iRunREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	return nil, false, nil
}

func (r *RerunS2iRunREST) New() runtime.Object {
	return &Rerun{}
}
