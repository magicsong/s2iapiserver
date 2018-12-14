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
	"fmt"
	"time"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/constants"

	"github.com/golang/glog"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
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
	BuilderName          string            `json:"builderName,omitempty"`
	BackoffLimit         int32             `json:"backoffLimit,omitempty"`
	SecondsAfterFinished int32             `json:"secondsAfterFinished,omitempty"`
	Environment          []EnvironmentSpec `json:"environment,omitempty"`
	OverideTag           string            `json:"overideTag,omitempty"`
}
type Payload struct {
	//TODO
}
type Trigger struct {
	Source  constants.TriggerSource `json:"source,omitempty"`
	Event   string                  `json:"event,omitempty"`
	Payload *Payload                `json:"payload,omitempty"`
}
type S2IRunResult struct {
	ImageName      string       `json:"imageName,omitempty"`
	CompletionTime *metav1.Time `json:"completionTime,omitempty" protobuf:"bytes,3,opt,name=completionTime"`
	Artifact       string       `json:"artifact,omitempty"`
}

// S2iRunStatus defines the observed state of S2iRun
type S2iRunStatus struct {
	StartTime *metav1.Time `json:"startTime,omitempty" protobuf:"bytes,2,opt,name=startTime"`

	RunState constants.RunningState `json:"runState,omitempty"`
	Trigger  *Trigger               `json:"trigger,omitempty"`
	Result   *S2IRunResult          `json:"result,omitempty"`
}

// Validate checks that an instance of S2iRun is well formed
func (S2iRunStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*devops.S2iRun)
	glog.V(4).Infof("Validating fields for S2iRun %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

func (S2iRunStrategy) ShortNames() []string {
	return []string{"s2ir"}
}

func (S2iRunStrategy) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1beta1.Table, error) {
	var table metav1beta1.Table
	var swaggerMetadataDescriptions = metav1.ObjectMeta{}.SwaggerDoc()
	table.ColumnDefinitions = []metav1beta1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Created At", Type: "date", Description: swaggerMetadataDescriptions["creationTimestamp"]},
		{Name: "ImageName", Type: "string", Description: "The name of output image"},
		{Name: "RunState", Type: "string", Description: "Current state of job"},
		{Name: "Triggered By", Type: "string", Description: "The source name who start the job"},
	}
	fn := func(obj runtime.Object) error {
		m, err := meta.Accessor(obj)
		if err != nil {
			return fmt.Errorf("the resource %s does not support being converted to a Table", "s2irun")
		}
		s := obj.(*devops.S2iRun)
		table.Rows = append(table.Rows, metav1beta1.TableRow{
			Cells: []interface{}{
				m.GetName(),
				m.GetCreationTimestamp().Time.UTC().Format(time.RFC3339),
				s.Status.Result.ImageName,
				s.Status.RunState,
				s.Status.Trigger.Source},
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
	return &table, nil
}

// DefaultingFunction sets default S2iRun field values
func (S2iRunSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*S2iRun)
	// set default field values here
	if obj.Status.Trigger == nil {
		t := new(Trigger)
		t.Source = constants.Default
		obj.Status.Trigger = t
	}
	if obj.Status.RunState == "" {
		obj.Status.RunState = constants.Unknown
	}
	if obj.Status.Result == nil {
		obj.Status.Result = new(S2IRunResult)
		obj.Status.Result.ImageName = "Here we are"
	}
	glog.V(4).Infof("Defaulting fields for S2iRun %s\n", obj.Name)
}
