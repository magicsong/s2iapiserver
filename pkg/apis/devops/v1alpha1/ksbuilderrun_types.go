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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KsBuilderRun
// +k8s:openapi-gen=true
// +resource:path=ksbuilderruns,strategy=KsBuilderRunStrategy
type KsBuilderRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KsBuilderRunSpec   `json:"spec,omitempty"`
	Status KsBuilderRunStatus `json:"status,omitempty"`
}

// KsBuilderRunSpec defines the desired state of KsBuilderRun
type KsBuilderRunSpec struct {
}

// KsBuilderRunStatus defines the observed state of KsBuilderRun
type KsBuilderRunStatus struct {
	StartTime *metav1.Time `json:"startTime,omitempty" protobuf:"bytes,2,opt,name=startTime"`

	// Represents time when the job was completed. It is not guaranteed to
	// be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty" protobuf:"bytes,3,opt,name=completionTime"`
	RunState       RunState     `json:"runState,omitempty"`
}

// Validate checks that an instance of KsBuilderRun is well formed
func (KsBuilderRunStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*devops.KsBuilderRun)
	log.Printf("Validating fields for KsBuilderRun %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default KsBuilderRun field values
func (KsBuilderRunSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*KsBuilderRun)
	// set default field values here
	log.Printf("Defaulting fields for KsBuilderRun %s\n", obj.Name)
}
