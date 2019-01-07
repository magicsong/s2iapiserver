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

	batchv1 "k8s.io/api/batch/v1"

	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	client "github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/clientset/typed/devops/v1alpha1"
	listers "github.com/magicsong/s2iapiserver/pkg/client/listers_generated/devops/v1alpha1"
	"github.com/magicsong/s2iapiserver/pkg/controller/sharedinformers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	batchv1listers "k8s.io/client-go/listers/batch/v1"
)

// +controller:group=devops,version=v1alpha1,kind=S2iRun,resource=s2iruns
type S2iRunControllerImpl struct {
	builders.DefaultControllerFns
	si     *sharedinformers.SharedInformers
	client *client.DevopsV1alpha1Client
	// lister indexes properties about S2iRun
	builderLister listers.S2iBuilderLister
	runLister     listers.S2iRunLister
	jobLister     batchv1listers.JobLister
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
	//c.jobLister = c.si.KubernetesFactory.Batch().V1().Jobs().Lister()
	c.AddWatches()
}

func defaultReconcileKey(s string) error {
	return nil
}
func (c *S2iRunControllerImpl) AddWatches() {
	//c.si.Watch("Jobs", c.si.KubernetesFactory.Batch().V1().Jobs().Informer(), c.JobToS2irun, defaultReconcileKey)
}

func (c *S2iRunControllerImpl) JobToS2irun(i interface{}) (string, error) {
	d, _ := i.(*batchv1.Job)
	if len(d.OwnerReferences) == 1 && d.OwnerReferences[0].Kind == "S2iRun" {
		return d.Namespace + "/" + d.OwnerReferences[0].Name, nil
	} else {
		// Not owned
		return "", nil
	}
}

// Reconcile handles enqueued messages
func (c *S2iRunControllerImpl) Reconcile(u *v1alpha1.S2iRun) error {
	// Implement controller logic here
	glog.V(2).Infof("Running reconcile S2iRun for %s\n", u.Name)
	instance := u.DeepCopy()
	err := c.Prepare(instance)
	if instance.Status.Result.CompletionTime != nil {
		return nil
	}
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
		_, err := c.client.S2iRuns(u.Namespace).UpdateStatus(instance)
		if err != nil {
			glog.Errorf("Upload instance <%s> of s2irun to storage failed,error:%s", u.Name, err.Error())
			return err
		}
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
			instance.Status.Result.ImageName = builder.Spec.Config.OutputImageName + "/" + builder.Spec.Config.Tag
		} else {
			instance.Status.Result.ImageName = builder.Spec.Config.OutputImageName + "/" + instance.Spec.OverideTag
		}
	}
	// job, err := c.jobLister.Jobs(instance.Namespace).Get(instance.Name + "-s2ijob")
	// if err != nil {
	// 	if !errors.IsNotFound(err) {
	// 		return err
	// 	}
	// 	instance.Status.RunState = constants.NotRunning
	// 	//create configmap
	// 	configMapName := instance.Name + "-s2iconfigmap"
	// 	_, err := c.si.KubernetesFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(instance.Namespace).Get(configMapName)
	// 	if err != nil {
	// 		if errors.IsNotFound(err) {
	// 			configmap, err := c.createConfigMap(instance, builder.Spec.Config)
	// 			if err != nil {
	// 				glog.Error("Error in generating configmap")
	// 				return err
	// 			}
	// 			err = controllerutil.SetControllerReference(instance, configmap, nil)
	// 			if err != nil {
	// 				glog.Errorln("Failed to SetControllerReference of configmap to s2irun")
	// 				return err
	// 			}
	// 			_, err = c.si.KubernetesClientSet.CoreV1().ConfigMaps(instance.Namespace).Create(configmap)
	// 			if err != nil {
	// 				glog.Error("Error in creating configmap in apiserver")
	// 				return err
	// 			}
	// 		}
	// 	}
	// 	//create job
	// 	job = c.createJob(instance)
	// 	_, err = c.si.KubernetesClientSet.BatchV1().Jobs(instance.Namespace).Create(job)
	// 	if err != nil {
	// 		glog.Error("Error in creating job in apiserver")
	// 		return err
	// 	}
	// } else {
	// 	if len(job.Status.Conditions) >= 1 {
	// 		instance.Status.Result.Message = job.Status.Conditions[0].Message
	// 	}
	// 	if job.Status.Active == 1 {
	// 		instance.Status.RunState = constants.Running
	// 	} else if job.Status.Failed == 1 {
	// 		instance.Status.RunState = constants.Failed
	// 	} else if job.Status.Succeeded == 1 {
	// 		instance.Status.RunState = constants.Successful
	// 		instance.Status.Result.CompletionTime = job.Status.CompletionTime
	// 	} else {
	// 		instance.Status.RunState = constants.Unknown
	// 	}
	// }
	return nil
}
