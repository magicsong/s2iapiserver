
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
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
)

// +controller:group=devops,version=v1alpha1,kind=KsBuilder,resource=ksbuilders
type KsBuilderControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about KsBuilder
	lister listers.KsBuilderLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *KsBuilderControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing ksbuilders labels
	c.lister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().KsBuilders().Lister()
}

// Reconcile handles enqueued messages
func (c *KsBuilderControllerImpl) Reconcile(u *v1alpha1.KsBuilder) error {
	// Implement controller logic here
	log.Printf("Running reconcile KsBuilder for %s\n", u.Name)
	return nil
}

func (c *KsBuilderControllerImpl) Get(namespace, name string) (*v1alpha1.KsBuilder, error) {
	return c.lister.KsBuilders(namespace).Get(name)
}
