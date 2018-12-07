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
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
)

// +controller:group=devops,version=v1alpha1,kind=S2iRun,resource=s2iruns
type S2iRunControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about S2iRun
	builderLister listers.S2iBuilderLister
	runLister     listers.S2iRunLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *S2iRunControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing s2iruns labels
	c.builderLister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().S2iBuilders().Lister()
	c.runLister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().S2iRuns().Lister()
}

// Reconcile handles enqueued messages
func (c *S2iRunControllerImpl) Reconcile(u *v1alpha1.S2iRun) error {
	// Implement controller logic here
	glog.V(2).Infof("Running reconcile S2iRun for %s\n", u.Name)
	instance, err := c.Get(u.Namespace, u.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	instance.Labels["builder"] = instance.Spec.BuilderName
	return nil
}

func (c *S2iRunControllerImpl) Get(namespace, name string) (*v1alpha1.S2iRun, error) {
	return c.runLister.S2iRuns(namespace).Get(name)
}
