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
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	jobListerv1 "k8s.io/client-go/listers/batch/v1"
	configmaplisterv1 "k8s.io/client-go/listers/core/v1"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
)

// +controller:group=devops,version=v1alpha1,kind=S2iBuilder,resource=s2ibuilders
type S2iBuilderControllerImpl struct {
	builders.DefaultControllerFns

<<<<<<< HEAD:pkg/controller/s2ibuilder/controller.go
	// lister indexes properties about S2iBuilder
	lister listers.S2iBuilderLister
=======
	// lister indexes properties about KsBuilderRun
	lister          listers.KsBuilderRunLister
	jobLister       jobListerv1.JobLister
	configMapLister configmaplisterv1.ConfigMapLister
}

func (c *KsBuilderControllerImpl) jobToKsBuilderRun(i interface{}) (string, error) {
	d := i.(*batchv1.Job)
	glog.V(2).Infof("Reconcile job <%s> belong to KsBuilderRun", d.Name)
	if len(d.OwnerReferences) == 1 && d.OwnerReferences[0].Kind == "KsBuilderRun" {
		return d.Namespace + "/" + d.OwnerReferences[0].Name, nil
	} else {
		// Not owned
		return "", nil
	}
}
func (c *KsBuilderRunControllerImpl) reconcileKey(key string) error {
	return nil
}

func (c *KsBuilderControllerImpl) configMapToKsBuilderRun(i interface{}) (string, error) {
	d := i.(*corev1.ConfigMap)
	glog.V(2).Infof("Reconcile configmap <%s> belong to KsBuilderRun", d.Name)
	if len(d.OwnerReferences) == 1 && d.OwnerReferences[0].Kind == "KsBuilderRun" {
		return d.Namespace + "/" + d.OwnerReferences[0].Name, nil
	} else {
		// Not owned
		return "", nil
	}
>>>>>>> 039424284d33a92560a99f054569465473f7e112:pkg/controller/ksbuilderrun/controller.go
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
<<<<<<< HEAD:pkg/controller/s2ibuilder/controller.go
func (c *S2iBuilderControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing s2ibuilders labels
	c.lister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().S2iBuilders().Lister()
=======
func (c *KsBuilderRunControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing ksbuilderruns labels
	c.lister = arguments.GetSharedInformers().Factory.Devops().V1alpha1().KsBuilderRuns().Lister()
	jobSI := arguments.GetSharedInformers().KubernetesFactory().Batch().V1().Jobs()
	c.jobLister = jobSI.Lister()
	arguments.GetSharedInformers().Watch("KsRunJob", jobSI, c.configMapToKsBuilderRun, c.reconcileKey)

	configmapSI := arguments.GetSharedInformers().KubernetesFactory().Core().V1().ConfigMaps()
	c.configMapLister = configmapSI.Lister()
	arguments.GetSharedInformers().Watch("KsRunConfigmap", configmapSI, c.configMapToKsBuilderRun, c.reconcileKey)
>>>>>>> 039424284d33a92560a99f054569465473f7e112:pkg/controller/ksbuilderrun/controller.go
}

// Reconcile handles enqueued messages
func (c *S2iBuilderControllerImpl) Reconcile(u *v1alpha1.S2iBuilder) error {
	// Implement controller logic here
<<<<<<< HEAD:pkg/controller/s2ibuilder/controller.go
	glog.V(1).Infof("Running reconcile S2iBuilder for %s\n", u.Name)
	glog.V(1).Infof("Hello,%s", u.Spec.Hello)
=======
	glog.V(1).Infof("Running reconcile KsBuilderRun for %s\n", u.Name)

>>>>>>> 039424284d33a92560a99f054569465473f7e112:pkg/controller/ksbuilderrun/controller.go
	return nil
}

func (c *S2iBuilderControllerImpl) Get(namespace, name string) (*v1alpha1.S2iBuilder, error) {
	return c.lister.S2iBuilders(namespace).Get(name)
}
