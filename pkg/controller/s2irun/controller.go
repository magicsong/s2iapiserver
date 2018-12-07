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
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	client "github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/clientset/typed/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
	"k8s.io/apimachinery/pkg/api/errors"
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
}

// Reconcile handles enqueued messages
func (c *S2iRunControllerImpl) Reconcile(u *v1alpha1.S2iRun) error {
	// Implement controller logic here
	glog.V(2).Infof("Running reconcile S2iRun for %s\n", u.Name)
	glog.Infof("Key:%+v", u)
	instance, err := c.Get(u.Namespace, u.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		glog.Errorf("Get instance <%s> of s2irun failed,error:%s", u.Name, err.Error())
		return err
	}
	instance.Labels["builder"] = instance.Spec.BuilderName
	if instance.Status.StartTime == nil {
		now := metav1.Now()
		instance.Status.StartTime = &now
	}
	instance, err = c.client.S2iRuns(u.Namespace).Update(instance)
	if err != nil {
		glog.Errorf("update instance <%s> of s2irun failed,error:%s", u.Name, err.Error())
		return err
	}
	return nil
}

func (c *S2iRunControllerImpl) Get(namespace, name string) (*v1alpha1.S2iRun, error) {
	return c.runLister.S2iRuns(namespace).Get(name)
}
