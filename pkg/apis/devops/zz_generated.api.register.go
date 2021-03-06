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

package devops

import (
	"context"
	"fmt"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	devopsconstants "github.com/magicsong/s2iapiserver/pkg/apis/devops/constants"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalS2iBuilder = builders.NewInternalResource(
		"s2ibuilders",
		"S2iBuilder",
		func() runtime.Object { return &S2iBuilder{} },
		func() runtime.Object { return &S2iBuilderList{} },
	)
	InternalS2iBuilderStatus = builders.NewInternalResourceStatus(
		"s2ibuilders",
		"S2iBuilderStatus",
		func() runtime.Object { return &S2iBuilder{} },
		func() runtime.Object { return &S2iBuilderList{} },
	)
	InternalS2iRun = builders.NewInternalResource(
		"s2iruns",
		"S2iRun",
		func() runtime.Object { return &S2iRun{} },
		func() runtime.Object { return &S2iRunList{} },
	)
	InternalS2iRunStatus = builders.NewInternalResourceStatus(
		"s2iruns",
		"S2iRunStatus",
		func() runtime.Object { return &S2iRun{} },
		func() runtime.Object { return &S2iRunList{} },
	)
	InternalRerunS2iRunREST = builders.NewInternalSubresource(
		"s2iruns", "Rerun", "rerun",
		func() runtime.Object { return &Rerun{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("devops.kubesphere.io").WithKinds(
		InternalS2iBuilder,
		InternalS2iBuilderStatus,
		InternalS2iRun,
		InternalS2iRunStatus,
		InternalRerunS2iRunREST,
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

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type S2iBuilder struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   S2iBuilderSpec
	Status S2iBuilderStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Rerun struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

type S2iBuilderStatus struct {
	RunCount     int
	LastRunState devopsconstants.RunningState
	LastRunName  *string
}

type S2iBuilderSpec struct {
	Config *S2iConfig
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type S2iRun struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   S2iRunSpec
	Status S2iRunStatus
}

type S2iConfig struct {
	DisplayName               string
	Description               string
	BuilderImage              string
	BuilderImageVersion       string
	BuilderBaseImageVersion   string
	RuntimeImage              string
	OutputImageName           string
	RuntimeImagePullPolicy    devopsconstants.PullPolicy
	RuntimeAuthentication     AuthConfig
	RuntimeArtifacts          []VolumeSpec
	DockerConfig              *DockerConfig
	PullAuthentication        AuthConfig
	PushAuthentication        AuthConfig
	IncrementalAuthentication AuthConfig
	DockerNetworkMode         devopsconstants.DockerNetworkMode
	PreserveWorkingDir        bool
	IgnoreSubmodules          bool
	Tag                       string
	BuilderPullPolicy         devopsconstants.PullPolicy
	PreviousImagePullPolicy   devopsconstants.PullPolicy
	Incremental               bool
	IncrementalFromTag        string
	RemovePreviousImage       bool
	Environment               []EnvironmentSpec
	LabelNamespace            string
	CallbackURL               string
	ScriptsURL                string
	Destination               string
	WorkingDir                string
	WorkingSourceDir          string
	LayeredBuild              bool
	Quiet                     bool
	ForceCopy                 bool
	ContextDir                string
	AssembleUser              string
	RunImage                  bool
	Usage                     bool
	Injections                []VolumeSpec
	CGroupLimits              *CGroupLimits
	DropCapabilities          []string
	ScriptDownloadProxyConfig *ProxyConfig
	ExcludeRegExp             string
	BlockOnBuild              bool
	HasOnBuild                bool
	BuildVolumes              []string
	Labels                    map[string]string
	SecurityOpt               []string
	KeepSymlinks              bool
	AsDockerfile              string
	ImageWorkDir              string
	ImageScriptsURL           string
	AddHost                   []string
	Export                    bool
	SourceURL                 string
}

type S2iRunStatus struct {
	StartTime *metav1.Time
	RunState  devopsconstants.RunningState
	Trigger   *Trigger
	Result    *S2IRunResult
}

type ProxyConfig struct {
	HTTPProxy  string
	HTTPSProxy string
}

type S2IRunResult struct {
	ImageName      string
	CompletionTime *metav1.Time
	Artifact       string
	Message        string
}

type Trigger struct {
	Source  devopsconstants.TriggerSource
	Event   string
	Payload *Payload
}

type CGroupLimits struct {
	MemoryLimitBytes int64
	CPUShares        int64
	CPUPeriod        int64
	CPUQuota         int64
	MemorySwap       int64
	Parent           string
}

type Payload struct {
}

type VolumeSpec struct {
	Source      string
	Destination string
	Keep        bool
}

type EnvironmentSpec struct {
	Name  string
	Value string
}

type AuthConfig struct {
	Username      string
	Password      string
	Email         string
	ServerAddress string
}

type DockerConfig struct {
	Endpoint  string
	CertFile  string
	KeyFile   string
	CAFile    string
	UseTLS    bool
	TLSVerify bool
}

type S2iRunSpec struct {
	BuilderName          string
	BackoffLimit         int32
	SecondsAfterFinished int32
	Environment          []EnvironmentSpec
	OverideTag           string
}

//
// S2iBuilder Functions and Structs
//
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
	metav1.TypeMeta
	metav1.ListMeta
	Items []S2iBuilder
}

func (S2iBuilder) NewStatus() interface{} {
	return S2iBuilderStatus{}
}

func (pc *S2iBuilder) GetStatus() interface{} {
	return pc.Status
}

func (pc *S2iBuilder) SetStatus(s interface{}) {
	pc.Status = s.(S2iBuilderStatus)
}

func (pc *S2iBuilder) GetSpec() interface{} {
	return pc.Spec
}

func (pc *S2iBuilder) SetSpec(s interface{}) {
	pc.Spec = s.(S2iBuilderSpec)
}

func (pc *S2iBuilder) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *S2iBuilder) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc S2iBuilder) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store S2iBuilder.
// +k8s:deepcopy-gen=false
type S2iBuilderRegistry interface {
	ListS2iBuilders(ctx context.Context, options *internalversion.ListOptions) (*S2iBuilderList, error)
	GetS2iBuilder(ctx context.Context, id string, options *metav1.GetOptions) (*S2iBuilder, error)
	CreateS2iBuilder(ctx context.Context, id *S2iBuilder) (*S2iBuilder, error)
	UpdateS2iBuilder(ctx context.Context, id *S2iBuilder) (*S2iBuilder, error)
	DeleteS2iBuilder(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewS2iBuilderRegistry(sp builders.StandardStorageProvider) S2iBuilderRegistry {
	return &storageS2iBuilder{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageS2iBuilder struct {
	builders.StandardStorageProvider
}

func (s *storageS2iBuilder) ListS2iBuilders(ctx context.Context, options *internalversion.ListOptions) (*S2iBuilderList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*S2iBuilderList), err
}

func (s *storageS2iBuilder) GetS2iBuilder(ctx context.Context, id string, options *metav1.GetOptions) (*S2iBuilder, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*S2iBuilder), nil
}

func (s *storageS2iBuilder) CreateS2iBuilder(ctx context.Context, object *S2iBuilder) (*S2iBuilder, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*S2iBuilder), nil
}

func (s *storageS2iBuilder) UpdateS2iBuilder(ctx context.Context, object *S2iBuilder) (*S2iBuilder, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*S2iBuilder), nil
}

func (s *storageS2iBuilder) DeleteS2iBuilder(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}

//
// S2iRun Functions and Structs
//
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
	metav1.TypeMeta
	metav1.ListMeta
	Items []S2iRun
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RerunList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Rerun
}

func (S2iRun) NewStatus() interface{} {
	return S2iRunStatus{}
}

func (pc *S2iRun) GetStatus() interface{} {
	return pc.Status
}

func (pc *S2iRun) SetStatus(s interface{}) {
	pc.Status = s.(S2iRunStatus)
}

func (pc *S2iRun) GetSpec() interface{} {
	return pc.Spec
}

func (pc *S2iRun) SetSpec(s interface{}) {
	pc.Spec = s.(S2iRunSpec)
}

func (pc *S2iRun) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *S2iRun) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc S2iRun) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store S2iRun.
// +k8s:deepcopy-gen=false
type S2iRunRegistry interface {
	ListS2iRuns(ctx context.Context, options *internalversion.ListOptions) (*S2iRunList, error)
	GetS2iRun(ctx context.Context, id string, options *metav1.GetOptions) (*S2iRun, error)
	CreateS2iRun(ctx context.Context, id *S2iRun) (*S2iRun, error)
	UpdateS2iRun(ctx context.Context, id *S2iRun) (*S2iRun, error)
	DeleteS2iRun(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewS2iRunRegistry(sp builders.StandardStorageProvider) S2iRunRegistry {
	return &storageS2iRun{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageS2iRun struct {
	builders.StandardStorageProvider
}

func (s *storageS2iRun) ListS2iRuns(ctx context.Context, options *internalversion.ListOptions) (*S2iRunList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*S2iRunList), err
}

func (s *storageS2iRun) GetS2iRun(ctx context.Context, id string, options *metav1.GetOptions) (*S2iRun, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*S2iRun), nil
}

func (s *storageS2iRun) CreateS2iRun(ctx context.Context, object *S2iRun) (*S2iRun, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*S2iRun), nil
}

func (s *storageS2iRun) UpdateS2iRun(ctx context.Context, object *S2iRun) (*S2iRun, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*S2iRun), nil
}

func (s *storageS2iRun) DeleteS2iRun(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}
