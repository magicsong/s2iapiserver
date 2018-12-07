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

package s2ibuilder

import (
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// +controller:group=devops,version=v1alpha1,kind=S2iBuilder,resource=s2ibuilders
type S2iBuilderControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about S2iBuilder
	builderLister listers.S2iBuilderLister
	runLister     listers.S2iRunLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *S2iBuilderControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing s2ibuilders labels
	c.builderLister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().S2iBuilders().Lister()
	c.runLister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().S2iRuns().Lister()
}

// Reconcile handles enqueued messages
func (c *S2iBuilderControllerImpl) Reconcile(u *v1alpha1.S2iBuilder) error {
	// Implement controller logic here
	glog.V(2).Infof("Running reconcile S2iBuilder for %s\n", u.Name)
	instance, err := c.Get(u.Namespace, u.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	runs, err := c.runLister.S2iRuns(u.Namespace).List(labels.Everything())
	if err != nil {
		glog.Errorf("cannot get s2irunners of s2ibuilder-<%s>,error is %s", u.Name, err.Error())
		return err
	}
	last := new(metav1.Time)
	for _, item := range runs {
		if item.Spec.BuilderName == u.Name {
			instance.Status.RunCount++
			if item.Status.StartTime != nil && item.Status.StartTime.After(last.Time) {
				last = item.Status.StartTime
				instance.Status.LastRunState = item.Status.RunState
				instance.Status.LastRunName = &(item.Name)
			}
		}
	}
	return nil
}

func (c *S2iBuilderControllerImpl) Get(namespace, name string) (*v1alpha1.S2iBuilder, error) {
	return c.builderLister.S2iBuilders(namespace).Get(name)
}
