
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

package ksbuilderrun

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
)

// +controller:group=devops,version=v1alpha1,kind=KsBuilderRun,resource=ksbuilderruns
type KsBuilderRunControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about KsBuilderRun
	lister listers.KsBuilderRunLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *KsBuilderRunControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing ksbuilderruns labels
	c.lister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().KsBuilderRuns().Lister()
}

// Reconcile handles enqueued messages
func (c *KsBuilderRunControllerImpl) Reconcile(u *v1alpha1.KsBuilderRun) error {
	// Implement controller logic here
	log.Printf("Running reconcile KsBuilderRun for %s\n", u.Name)
	return nil
}

func (c *KsBuilderRunControllerImpl) Get(namespace, name string) (*v1alpha1.KsBuilderRun, error) {
	return c.lister.KsBuilderRuns(namespace).Get(name)
}
