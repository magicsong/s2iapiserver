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
	"log"

	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// S2iBuilder
// +k8s:openapi-gen=true
// +resource:path=s2ibuilders,strategy=S2iBuilderStrategy
type S2iBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S2iBuilderSpec   `json:"spec,omitempty"`
	Status S2iBuilderStatus `json:"status,omitempty"`
}

// S2iBuilderSpec defines the desired state of S2iBuilder
type S2iBuilderSpec struct {
	Config string `json:"config,omitempty"`
}

const (
	NotRunning string = "Not Running Yet"
	Successful        = "Successful"
	Failed            = "Failed"
	Unknown           = "Unknown"
)

// S2iBuilderStatus defines the observed state of S2iBuilder
type S2iBuilderStatus struct {
	RunCount     int    `json:"runCount,omitempty"`
	LastRunState string `json:"lastRunState,omitempty"`
	LastRunName  string `json:"lastRunName,omitempty"`
}

// Validate checks that an instance of S2iBuilder is well formed
func (S2iBuilderStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*devops.S2iBuilder)
	log.Printf("Validating fields for S2iBuilder %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default S2iBuilder field values
func (S2iBuilderSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*S2iBuilder)
	obj.Status.LastRunState = NotRunning
}
