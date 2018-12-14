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

package s2irun

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	client "github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/clientset/typed/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +controller:group=devops,version=v1alpha1,kind=S2iRun,resource=s2iruns
type S2iRunControllerImpl struct {
	builders.DefaultControllerFns
	si     *sharedinformers.SharedInformers
	client *client.DevopsV1alpha1Client
	// lister indexes properties about S2iRun
	builderLister listers.S2iBuilderLister
	runLister     listers.S2iRunLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *S2iRunControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing s2iruns labels
	c.si = arguments.GetSharedInformers()
	api := c.si.Factory.Devops().V1alpha1()
	c.client = client.NewForConfigOrDie(arguments.GetRestConfig())
	c.builderLister = api.S2iBuilders().Lister()
	c.runLister = api.S2iRuns().Lister()
	c.AddWatches()
}

func (c *S2iRunControllerImpl) AddWatches() {
	c.si.Watch("Builders", c.si.Factory.Devops().V1alpha1().S2iBuilders().Informer(), func(i interface{}) (string, error) {
		d, _ := i.(*v1alpha1.S2iBuilder)
		glog.Infoln("Watched builder")
		return d.Namespace + "/" + d.Name, nil
	}, func(s string) error {
		return nil
	})
}

// Reconcile handles enqueued messages
func (c *S2iRunControllerImpl) Reconcile(u *v1alpha1.S2iRun) error {
	// Implement controller logic here
	glog.V(2).Infof("Running reconcile S2iRun for %s\n", u.Name)
	instance := u.DeepCopy()
	err := c.Prepare(instance)
	if err != nil {
		glog.Errorf("Failed to preprocess s2run instance <%s>, error:%s", u.Name, err.Error())
		return err
	}
	err = c.UpdateStatus(instance)
	if err != nil {
		glog.Errorf("Failed to update s2run instance <%s>, error:%s", u.Name, err.Error())
		return err
	}
	if !reflect.DeepEqual(u.Status, instance.Status) {
		instance, err := c.client.S2iRuns(u.Namespace).UpdateStatus(instance)
		if err != nil {
			glog.Errorf("Upload instance <%s> of s2irun to storage failed,error:%s", u.Name, err.Error())
			return err
		}
		glog.Infof("%+v", instance)
	}
	return nil
}

func (c *S2iRunControllerImpl) Get(namespace, name string) (*v1alpha1.S2iRun, error) {
	return c.runLister.S2iRuns(namespace).Get(name)
}
func (c *S2iRunControllerImpl) Prepare(instance *v1alpha1.S2iRun) error {
	if instance.Labels == nil {
		instance.Labels = make(map[string]string)
	}
	instance.Labels["builder"] = instance.Spec.BuilderName
	return nil
}
func (c *S2iRunControllerImpl) UpdateStatus(instance *v1alpha1.S2iRun) error {
	if instance.Status.StartTime == nil {
		now := metav1.Now()
		instance.Status.StartTime = &now
	}
	builder, err := c.builderLister.S2iBuilders(instance.Namespace).Get(instance.Spec.BuilderName)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
	} else {
		if instance.Spec.OverideTag == "" {
			instance.Status.Result.ImageName = builder.Spec.OutputImageName + "/" + builder.Spec.Tag
		} else {
			instance.Status.Result.ImageName = builder.Spec.OutputImageName + "/" + instance.Spec.OverideTag
		}
	}

	return nil
}
