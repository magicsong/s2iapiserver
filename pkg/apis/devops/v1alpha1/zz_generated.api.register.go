/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	devopsS2iBuilderStorage = builders.NewApiResource( // Resource status endpoint
		devops.InternalS2iBuilder,
		S2iBuilderSchemeFns{},
		func() runtime.Object { return &S2iBuilder{} },     // Register versioned resource
		func() runtime.Object { return &S2iBuilderList{} }, // Register versioned resource list
		&S2iBuilderStrategy{builders.StorageStrategySingleton},
	)
	devopsS2iRunStorage = builders.NewApiResource( // Resource status endpoint
		devops.InternalS2iRun,
		S2iRunSchemeFns{},
		func() runtime.Object { return &S2iRun{} },     // Register versioned resource
		func() runtime.Object { return &S2iRunList{} }, // Register versioned resource list
		&S2iRunStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("devops.kubesphere.io", "v1alpha1").WithResources(
		devopsS2iBuilderStorage,
		builders.NewApiResource( // Resource status endpoint
			devops.InternalS2iBuilderStatus,
			S2iBuilderSchemeFns{},
			func() runtime.Object { return &S2iBuilder{} },     // Register versioned resource
			func() runtime.Object { return &S2iBuilderList{} }, // Register versioned resource list
			&S2iBuilderStatusStrategy{builders.StatusStorageStrategySingleton},
		), devopsS2iRunStorage,
		builders.NewApiResource( // Resource status endpoint
			devops.InternalS2iRunStatus,
			S2iRunSchemeFns{},
			func() runtime.Object { return &S2iRun{} },     // Register versioned resource
			func() runtime.Object { return &S2iRunList{} }, // Register versioned resource list
			&S2iRunStatusStrategy{builders.StatusStorageStrategySingleton},
		), builders.NewApiResourceWithStorage(
			devops.InternalRerunS2iRunREST,
			builders.SchemeFnsSingleton,
			func() runtime.Object { return &Rerun{} }, // Register versioned resource
			nil,
			func() rest.Storage { return &RerunS2iRunREST{devops.NewS2iRunRegistry(devopsS2iRunStorage)} },
		),
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//
// S2iBuilder Functions and Structs
//
// +k8s:deepcopy-gen=false
type S2iBuilderSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type S2iBuilderStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type S2iBuilderStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type S2iBuilderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S2iBuilder `json:"items"`
}

//
// S2iRun Functions and Structs
//
// +k8s:deepcopy-gen=false
type S2iRunSchemeFns struct {
	builders.DefaultSchemeFns
}

// +k8s:deepcopy-gen=false
type S2iRunStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type S2iRunStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type S2iRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S2iRun `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RerunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rerun `json:"items"`
}
