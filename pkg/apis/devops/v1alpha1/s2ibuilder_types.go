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

package v1alpha1

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/magicsong/s2iapiserver/pkg/apis/devops"
	"github.com/magicsong/s2iapiserver/pkg/apis/devops/constants"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// S2iBuilder
// +k8s:openapi-gen=true
// +resource:path=s2ibuilders,strategy=S2iBuilderStrategy
type S2iBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S2iBuilderSpec   `json:"spec,omitempty"`
	Status S2iBuilderStatus `json:"status,omitempty"`
}

// S2iBuilderSpec defines the desired state of S2iBuilder
type S2iBuilderSpec struct {
	// DisplayName is a result image display-name label. This defaults to the
	// output image name.
	DisplayName string `json:"displayName,omitempty"`

	// Description is a result image description label. The default is no
	// description.
	Description string `json:"description,omitempty"`

	// BuilderImage describes which image is used for building the result images.
	BuilderImage string `json:"builderImage,omitempty"`

	// BuilderImageVersion provides optional version information about the builder image.
	BuilderImageVersion string `json:"builderImageVersion,omitempty"`

	// BuilderBaseImageVersion provides optional version information about the builder base image.
	BuilderBaseImageVersion string `json:"builderBaseImageVersion,omitempty"`

	// RuntimeImage specifies the image that will be a base for resulting image
	// and will be used for running an application. By default, BuilderImage is
	// used for building and running, but the latter may be overridden.
	RuntimeImage string `json:"runtimeImage,omitempty"`

	// RuntimeImagePullPolicy specifies when to pull a runtime image.
	RuntimeImagePullPolicy constants.PullPolicy `json:"runtimeImagePullPolicy,omitempty"`

	// RuntimeAuthentication holds the authentication information for pulling the
	// runtime Docker images from private repositories.
	RuntimeAuthentication AuthConfig `json:"runtimeAuthentication,omitempty"`

	// RuntimeArtifacts specifies a list of source/destination pairs that will
	// be copied from builder to a runtime image. Source can be a file or
	// directory. Destination must be a directory. Regardless whether it
	// is an absolute or relative path, it will be placed into image's WORKDIR.
	// Destination also can be empty or equals to ".", in this case it just
	// refers to a root of WORKDIR.
	// In case it's empty, S2I will try to get this list from
	// io.openshift.s2i.assemble-input-files label on a RuntimeImage.
	RuntimeArtifacts []VolumeSpec `json:"runtimeArtifacts,omitempty"`

	// DockerConfig describes how to access host docker daemon.
	DockerConfig *DockerConfig `json:"dockerConfig,omitempty"`

	// PullAuthentication holds the authentication information for pulling the
	// Docker images from private repositories
	PullAuthentication AuthConfig `json:"pullAuthentication,omitempty"`

	// PullAuthentication holds the authentication information for pulling the
	// Docker images from private repositories
	PushAuthentication AuthConfig `json:"pushAuthentication,omitempty"`

	// IncrementalAuthentication holds the authentication information for pulling the
	// previous image from private repositories
	IncrementalAuthentication AuthConfig `json:"incrementalAuthentication,omitempty"`

	// DockerNetworkMode is used to set the docker network setting to --net=container:<id>
	// when the builder is invoked from a container.
	DockerNetworkMode constants.DockerNetworkMode `json:"dockerNetworkMode,omitempty"`

	// PreserveWorkingDir describes if working directory should be left after processing.
	PreserveWorkingDir bool `json:"preserveWorkingDir,omitempty"`

	// IgnoreSubmodules determines whether we will attempt to pull in submodules
	// (via --recursive or submodule init)
	IgnoreSubmodules bool `json:"ignoreSubmodules,omitempty"`

	// Tag is a result image tag name.
	Tag string `json:"tag,omitempty"`

	// BuilderPullPolicy specifies when to pull the builder image
	BuilderPullPolicy constants.PullPolicy `json:"builderPullPolicy,omitempty"`

	// PreviousImagePullPolicy specifies when to pull the previously build image
	// when doing incremental build
	PreviousImagePullPolicy constants.PullPolicy `json:"previousImagePullPolicy,omitempty"`

	// Incremental describes whether to try to perform incremental build.
	Incremental bool `json:"incremental,omitempty"`

	// IncrementalFromTag sets an alternative image tag to look for existing
	// artifacts. Tag is used by default if this is not set.
	IncrementalFromTag string `json:"incrementalFromTag,omitempty"`

	// RemovePreviousImage describes if previous image should be removed after successful build.
	// This applies only to incremental builds.
	RemovePreviousImage bool `json:"removePreviousImage,omitempty"`

	// Environment is a map of environment variables to be passed to the image.
	Environment []EnvironmentSpec `json:"environment,omitempty"`

	// LabelNamespace provides the namespace under which the labels will be generated.
	LabelNamespace string `json:"labelNamespace,omitempty"`

	// CallbackURL is a URL which is called upon successful build to inform about that fact.
	CallbackURL string `json:"callbackUrl,omitempty"`

	// ScriptsURL is a URL describing where to fetch the S2I scripts from during build process.
	// This url can be a reference within the builder image if the scheme is specified as image://
	ScriptsURL string `json:"scriptsUrl,omitempty"`

	// Destination specifies a location where the untar operation will place its artifacts.
	Destination string `json:"destination,omitempty"`

	// WorkingDir describes temporary directory used for downloading sources, scripts and tar operations.
	WorkingDir string `json:"workingDir,omitempty"`

	// WorkingSourceDir describes the subdirectory off of WorkingDir set up during the repo download
	// that is later used as the root for ignore processing
	WorkingSourceDir string `json:"workingSourceDir,omitempty"`

	// LayeredBuild describes if this is build which layered scripts and sources on top of BuilderImage.
	LayeredBuild bool `json:"layeredBuild,omitempty"`

	// Operate quietly. Progress and assemble script output are not reported, only fatal errors.
	// (default: false).
	Quiet bool `json:"quiet,omitempty"`

	// ForceCopy results in only the file SCM plugin being used (i.e. no `git clone`); allows for empty directories to be included
	// in resulting image (since git does not support that).
	// (default: false).
	ForceCopy bool `json:"forceCopy,omitempty"`

	// Specify a relative directory inside the application repository that should
	// be used as a root directory for the application.
	ContextDir string `json:"contextDir,omitempty"`

	// AssembleUser specifies the user to run the assemble script in container
	AssembleUser string `json:"assembleUser,omitempty"`

	// RunImage will trigger a "docker run ..." invocation of the produced image so the user
	// can see if it operates as he would expect
	RunImage bool `json:"runImage,omitempty"`

	// Usage allows for properly shortcircuiting s2i logic when `s2i usage` is invoked
	Usage bool `json:"usage,omitempty"`

	// Injections specifies a list source/destination folders that are injected to
	// the container that runs assemble.
	// All files we inject will be truncated after the assemble script finishes.
	Injections []VolumeSpec `json:"injections,omitempty"`

	// CGroupLimits describes the cgroups limits that will be applied to any containers
	// run by s2i.
	CGroupLimits *CGroupLimits `json:"cgroupLimits,omitempty"`

	// DropCapabilities contains a list of capabilities to drop when executing containers
	DropCapabilities []string `json:"dropCapabilities,omitempty"`

	// ScriptDownloadProxyConfig optionally specifies the http and https proxy
	// to use when downloading scripts
	ScriptDownloadProxyConfig *ProxyConfig `json:"scriptDownloadProxyConfig,omitempty"`

	// ExcludeRegExp contains a string representation of the regular expression desired for
	// deciding which files to exclude from the tar stream
	ExcludeRegExp string `json:"excludeRegExp,omitempty"`

	// BlockOnBuild prevents s2i from performing a docker build operation
	// if one is necessary to execute ONBUILD commands, or to layer source code into
	// the container for images that don't have a tar binary available, if the
	// image contains ONBUILD commands that would be executed.
	BlockOnBuild bool `json:"blockOnBuild,omitempty"`

	// HasOnBuild will be set to true if the builder image contains ONBUILD instructions
	HasOnBuild bool `json:"hasOnBuild,omitempty"`

	// BuildVolumes specifies a list of volumes to mount to container running the
	// build.
	BuildVolumes []string `json:"buildVolumes,omitempty"`

	// Labels specify labels and their values to be applied to the resulting image. Label keys
	// must have non-zero length. The labels defined here override generated labels in case
	// they have the same name.
	Labels map[string]string `json:"labels,omitempty"`

	// SecurityOpt are passed as options to the docker containers launched by s2i.
	SecurityOpt []string `json:"securityOpt,omitempty"`

	// KeepSymlinks indicates to copy symlinks as symlinks. Default behavior is to follow
	// symlinks and copy files by content.
	KeepSymlinks bool `json:"keepSymlinks,omitempty"`

	// AsDockerfile indicates the path where the Dockerfile should be written instead of building
	// a new image.
	AsDockerfile string `json:"asDockerfile,omitempty"`

	// ImageWorkDir is the default working directory for the builder image.
	ImageWorkDir string `json:"imageWorkDir,omitempty"`

	// ImageScriptsURL is the default location to find the assemble/run scripts for a builder image.
	// This url can be a reference within the builder image if the scheme is specified as image://
	ImageScriptsURL string `json:"imageScriptsUrl,omitempty"`

	// AddHost Add a line to /etc/hosts for test purpose or private use in LAN. Its format is host:IP,muliple hosts can be added  by using multiple --add-host
	AddHost []string `json:"addHost,omitempty"`

	//Export Push the result image to specify image registry in tag
	Export bool `json:"export,omitempty"`

	SourceURL string `json:"sourceUrl,omitempty"`
}

// S2iBuilderStatus defines the observed state of S2iBuilder
type S2iBuilderStatus struct {
	RunCount     int                    `json:"runCount,omitempty"`
	LastRunState constants.RunningState `json:"lastRunState,omitempty"`
	LastRunName  *string                `json:"lastRunName,omitempty"`
}

// Validate checks that an instance of S2iBuilder is well formed
func (S2iBuilderStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*devops.S2iBuilder)
	log.Printf("Validating fields for S2iBuilder %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default S2iBuilder field values
func (S2iBuilderSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*S2iBuilder)
	obj.Status.LastRunState = constants.NotRunning
}

func (S2iBuilderStrategy) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1beta1.Table, error) {
	var table metav1beta1.Table
	var swaggerMetadataDescriptions = metav1.ObjectMeta{}.SwaggerDoc()
	table.ColumnDefinitions = []metav1beta1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Created At", Type: "date", Description: swaggerMetadataDescriptions["creationTimestamp"]},
		{Name: "RunCount", Type: "interger", Description: "The aggregate readiness state of this pod for accepting traffic."},
		{Name: "RunState", Type: "string", Description: "The aggregate status of the containers in this pod."},
	}
	fn := func(obj runtime.Object) error {
		m, err := meta.Accessor(obj)
		if err != nil {
			return fmt.Errorf("the resource %s does not support being converted to a Table", "s2ibuilder")
		}
		s := obj.(*S2iBuilder)
		table.Rows = append(table.Rows, metav1beta1.TableRow{
			Cells: []interface{}{
				m.GetName(),
				m.GetCreationTimestamp().Time.UTC().Format(time.RFC3339),
				s.Status.RunCount,
				s.Status.LastRunState},
			Object: runtime.RawExtension{Object: obj},
		})
		return nil
	}
	switch {
	case meta.IsListType(object):
		if err := meta.EachListItem(object, fn); err != nil {
			return nil, err
		}
	default:
		if err := fn(object); err != nil {
			return nil, err
		}
	}
	if m, err := meta.ListAccessor(object); err == nil {
		table.ResourceVersion = m.GetResourceVersion()
		table.SelfLink = m.GetSelfLink()
		table.Continue = m.GetContinue()
	} else {
		if m, err := meta.CommonAccessor(object); err == nil {
			table.ResourceVersion = m.GetResourceVersion()
			table.SelfLink = m.GetSelfLink()
		}
	}
	return &table, nil
}
