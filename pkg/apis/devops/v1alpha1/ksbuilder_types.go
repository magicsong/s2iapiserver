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
	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops"
	"github.com/magicsong/s2irun/pkg/api"
)

type RunState string

const (
	NotRunning RunState = "Not Running Yet"
	Successful          = "Successful"
	Failed              = "Failed"
	Unknown             = "Unknown"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KsBuilder
// +k8s:openapi-gen=true
// +resource:path=ksbuilders,strategy=KsBuilderStrategy
type KsBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KsBuilderSpec   `json:"spec,omitempty"`
	Status KsBuilderStatus `json:"status,omitempty"`
}

// KsBuilderSpec defines the desired state of KsBuilder
type KsBuilderSpec struct {
	Config api.Config `json:"config,omitempty"`
}

// KsBuilderStatus defines the observed state of KsBuilder
type KsBuilderStatus struct {
	RunCount     int      `json:"runCount,omitempty"`
	LastRunState RunState `json:"lastRunState,omitempty"`
	LastRunName  string   `json:"lastRunName,omitempty"`
}

// Validate checks that an instance of KsBuilder is well formed
func (KsBuilderStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*devops.KsBuilder)
	glog.V(2).Infof("Validating fields for KsBuilder %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default KsBuilder field values
func (KsBuilderSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*KsBuilder)
	// set default field values here
	obj.Status.LastRunState = NotRunning
}
