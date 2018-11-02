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

package ksbuilder

import (
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	batchv1 "k8s.io/api/batch/v1"
	joblister "k8s.io/client-go/listers/batch/v1"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
)

// +controller:group=devops,version=v1alpha1,kind=KsBuilder,resource=ksbuilders
type KsBuilderControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about KsBuilder
	lister    listers.KsBuilderLister
	jobLister joblister.JobLister
}

func (c *KsBuilderControllerImpl) jobToKsBuilder(i interface{}) (string, error) {
	d := i.(*batchv1.Job)
	glog.V(2).Infof("Reconcile job <%s> belong to KsBuilder", d.Name)
	if len(d.OwnerReferences) == 1 && d.OwnerReferences[0].Kind == "KsBuilder" {
		return d.Namespace + "/" + d.OwnerReferences[0].Name, nil
	} else {
		// Not owned
		return "", nil
	}
}

func (c *KsBuilderControllerImpl) reconcileKey(key string) error {
	return nil
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *KsBuilderControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing ksbuilders labels
	c.lister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().KsBuilders().Lister()
	jobSi := arguments.GetSharedInformers().KubernetesFactory().Batch().V1().Jobs()
	arguments.GetSharedInformers().Watch("KsBuilderJob", jobSi, c.jobToKsBuilder, c.reconcileKey)

}

// Reconcile handles enqueued messages
func (c *KsBuilderControllerImpl) Reconcile(u *v1alpha1.KsBuilder) error {
	// Implement controller logic here
	glog.V(2).Infof("Running reconcile KsBuilder for %s", u.Name)
	instance, err := c.Get(u.Namespace, u.Name)
	if err != nil {
		glog.Errorf("Get KsBuilder %s failed,error %v:", u.Namespace+"/"+u.Name, err)
		return err
	}

	return nil
}

func (c *KsBuilderControllerImpl) Get(namespace, name string) (*v1alpha1.KsBuilder, error) {
	return c.lister.KsBuilders(namespace).Get(name)
}
