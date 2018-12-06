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

// S2iRun
// +k8s:openapi-gen=true
// +resource:path=s2iruns,strategy=S2iRunStrategy
type S2iRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S2iRunSpec   `json:"spec,omitempty"`
	Status S2iRunStatus `json:"status,omitempty"`
}

// S2iRunSpec defines the desired state of S2iRun
type S2iRunSpec struct {
	BuilderName          string `json:"builderName,omitempty"`
	BackoffLimit         int32  `json:"backoffLimit,omitempty"`
	SecondsAfterFinished int32  `json:"secondsAfterFinished,omitempty"`
}

// S2iRunStatus defines the observed state of S2iRun
type S2iRunStatus struct {
	StartTime      *metav1.Time `json:"startTime,omitempty" protobuf:"bytes,2,opt,name=startTime"`
	CompletionTime *metav1.Time `json:"completionTime,omitempty" protobuf:"bytes,3,opt,name=completionTime"`
	RunState       string       `json:"runState,omitempty"`
}

// Validate checks that an instance of S2iRun is well formed
func (S2iRunStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*devops.S2iRun)
	log.Printf("Validating fields for S2iRun %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default S2iRun field values
func (S2iRunSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*S2iRun)
	// set default field values here
	log.Printf("Defaulting fields for S2iRun %s\n", obj.Name)
}
