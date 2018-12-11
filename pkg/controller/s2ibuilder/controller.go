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
	"reflect"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	client "github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/clientset/typed/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// +controller:group=devops,version=v1alpha1,kind=S2iBuilder,resource=s2ibuilders
type S2iBuilderControllerImpl struct {
	builders.DefaultControllerFns
	si *sharedinformers.SharedInformers
	// lister indexes properties about S2iBuilder
	builderLister listers.S2iBuilderLister
	runLister     listers.S2iRunLister
	client        *client.DevopsV1alpha1Client
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *S2iBuilderControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing s2ibuilders labels
	c.si = arguments.GetSharedInformers()
	c.builderLister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().S2iBuilders().Lister()
	c.runLister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().S2iRuns().Lister()
	c.client = client.NewForConfigOrDie(arguments.GetRestConfig())
	c.si.Watch("BuilderRun", c.si.Factory.Devops().V1alpha1().S2iRuns().Informer(), func(i interface{}) (string, error) {
		d, _ := i.(*v1alpha1.S2iRun)
		glog.V(1).Infof("[s2ibuilder] Reconcile key for s2irun")
		return d.Namespace + "/" + d.Name, nil
	}, func(s string) error {
		return nil
	})
}

// Reconcile handles enqueued messages
func (c *S2iBuilderControllerImpl) Reconcile(u *v1alpha1.S2iBuilder) error {
	// Implement controller logic here
	glog.V(1).Infof("Running reconcile S2iBuilder for %s\n", u.Name)
	instance := u.DeepCopy()
	instance.Status = v1alpha1.S2iBuilderStatus{}
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
	//glog.V(1).Infof("Status origin %d,Status change %d", u.Status.RunCount, instance.Status.RunCount)
	if !reflect.DeepEqual(u.Status, instance.Status) {
		_, err = c.client.S2iBuilders(u.Namespace).UpdateStatus(instance)
		if err != nil {
			glog.Errorf("cannot update s2ibuilder-<%s>,error is %s", u.Name, err.Error())
			return err
		}
	}
	return nil
}

func (c *S2iBuilderControllerImpl) Get(namespace, name string) (*v1alpha1.S2iBuilder, error) {
	return c.builderLister.S2iBuilders(namespace).Get(name)
}
