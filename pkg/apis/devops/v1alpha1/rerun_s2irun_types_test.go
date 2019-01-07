
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

package v1alpha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/magicsong/s2iapiserver/pkg/apis/devops/v1alpha1"
	. "github.com/magicsong/s2iapiserver/pkg/client/clientset_generated/clientset/typed/devops/v1alpha1"
)

var _ = Describe("S2iRun", func() {
	var instance S2iRun
	var expected S2iRun
	var client S2iRunInterface

	BeforeEach(func() {
		instance = S2iRun{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	Describe("when sending a rerun request", func() {
		It("should return success", func() {
			client = cs.DevopsV1alpha1Client.S2iruns("s2irun-test-rerun")
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			rerun := &Rerun{}
			rerun.Name = instance.Name
			restClient := cs.DevopsV1alpha1Client.RESTClient()
			err = restClient.Post().Namespace("s2irun-test-rerun").
				Name(instance.Name).
				Resource("s2iruns").
				SubResource("rerun").
				Body(rerun).Do().Error()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
