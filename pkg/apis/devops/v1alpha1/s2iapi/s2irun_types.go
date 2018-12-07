package s2iapi

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

// invalidFilenameCharacters contains a list of character we consider malicious
// when injecting the directories into containers.
const invalidFilenameCharacters = `;*?"<>|%#$!+{}&[],"'` + "`"

const (
	// PullAlways means that we always attempt to pull the latest image.
	PullAlways PullPolicy = "always"

	// PullNever means that we never pull an image, but only use a local image.
	PullNever PullPolicy = "never"

	// PullIfNotPresent means that we pull if the image isn't present on disk.
	PullIfNotPresent PullPolicy = "if-not-present"

	// DefaultBuilderPullPolicy specifies the default pull policy to use
	DefaultBuilderPullPolicy = PullIfNotPresent

	// DefaultRuntimeImagePullPolicy specifies the default pull policy to use.
	DefaultRuntimeImagePullPolicy = PullIfNotPresent

	// DefaultPreviousImagePullPolicy specifies policy for pulling the previously
	// build Docker image when doing incremental build
	DefaultPreviousImagePullPolicy = PullIfNotPresent
)

// Config contains essential fields for performing build.
//+k8s:openapi-gen=true
type Config struct {
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
	RuntimeImagePullPolicy PullPolicy `json:"runtimeImagePullPolicy,omitempty"`

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
	RuntimeArtifacts VolumeList `json:"runtimeArtifacts,omitempty"`

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
	DockerNetworkMode DockerNetworkMode `json:"dockerNetworkMode,omitempty"`

	// PreserveWorkingDir describes if working directory should be left after processing.
	PreserveWorkingDir bool `json:"preserveWorkingDir,omitempty"`

	// IgnoreSubmodules determines whether we will attempt to pull in submodules
	// (via --recursive or submodule init)
	IgnoreSubmodules bool `json:"ignoreSubmodules,omitempty"`

	// Tag is a result image tag name.
	Tag string `json:"tag,omitempty"`

	// BuilderPullPolicy specifies when to pull the builder image
	BuilderPullPolicy PullPolicy `json:"builderPullPolicy,omitempty"`

	// PreviousImagePullPolicy specifies when to pull the previously build image
	// when doing incremental build
	PreviousImagePullPolicy PullPolicy `json:"previousImagePullPolicy,omitempty"`

	// Incremental describes whether to try to perform incremental build.
	Incremental bool `json:"incremental,omitempty"`

	// IncrementalFromTag sets an alternative image tag to look for existing
	// artifacts. Tag is used by default if this is not set.
	IncrementalFromTag string `json:"incrementalFromTag,omitempty"`

	// RemovePreviousImage describes if previous image should be removed after successful build.
	// This applies only to incremental builds.
	RemovePreviousImage bool `json:"removePreviousImage,omitempty"`

	// Environment is a map of environment variables to be passed to the image.
	Environment EnvironmentList `json:"environment,omitempty"`

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
	Injections VolumeList `json:"injections,omitempty"`

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

// DeepCopyInto to implement k8s api requirement
func (c *Config) DeepCopyInto(out *Config) {
	*out = *c

	//slice
	if c.DropCapabilities != nil {
		out.DropCapabilities = make([]string, len(c.DropCapabilities))
		copy(out.DropCapabilities, c.DropCapabilities)
	}
	if c.BuildVolumes != nil {
		out.BuildVolumes = make([]string, len(c.BuildVolumes))
		copy(out.BuildVolumes, c.BuildVolumes)
	}
	if c.AddHost != nil {
		out.AddHost = make([]string, len(c.AddHost))
		copy(out.AddHost, c.AddHost)
	}
	if c.SecurityOpt != nil {
		out.SecurityOpt = make([]string, len(c.SecurityOpt))
		copy(out.SecurityOpt, c.SecurityOpt)
	}

	//pointer
	if c.DockerConfig != nil {
		out.DockerConfig = new(DockerConfig)
		*(out.DockerConfig) = *(c.DockerConfig)
	}
	if c.CGroupLimits != nil {
		out.CGroupLimits = new(CGroupLimits)
		*(out.CGroupLimits) = *(c.CGroupLimits)
	}
}

// EnvironmentSpec specifies a single environment variable.
type EnvironmentSpec struct {
	Name  string
	Value string
}

// EnvironmentList contains list of environment variables.
type EnvironmentList []EnvironmentSpec

// ProxyConfig holds proxy configuration.
type ProxyConfig struct {
	HTTPProxy  *url.URL
	HTTPSProxy *url.URL
}

// CGroupLimits holds limits used to constrain container resources.
type CGroupLimits struct {
	MemoryLimitBytes int64
	CPUShares        int64
	CPUPeriod        int64
	CPUQuota         int64
	MemorySwap       int64
	Parent           string
}

// VolumeSpec represents a single volume mount point.
type VolumeSpec struct {
	// Source is a reference to the volume source.
	Source string
	// Destination is the path to mount the volume to - absolute or relative.
	Destination string
	// Keep indicates if the mounted data should be kept in the final image.
	Keep bool
}

// VolumeList contains list of VolumeSpec.
type VolumeList []VolumeSpec

// DockerConfig contains the configuration for a Docker connection.
type DockerConfig struct {
	// Endpoint is the docker network endpoint or socket
	Endpoint string

	// CertFile is the certificate file path for a TLS connection
	CertFile string

	// KeyFile is the key file path for a TLS connection
	KeyFile string

	// CAFile is the certificate authority file path for a TLS connection
	CAFile string

	// UseTLS indicates if TLS must be used
	UseTLS bool

	// TLSVerify indicates if TLS peer must be verified
	TLSVerify bool
}

// AuthConfig is our abstraction of the Registry authorization information for whatever
// docker client we happen to be based on
type AuthConfig struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email,omitempty"`
	ServerAddress string `json:"server_address,omitempty"`
}

// ContainerConfig is the abstraction of the docker client provider (formerly go-dockerclient, now either
// engine-api or kube docker client) container.Config type that is leveraged by s2i or origin
type ContainerConfig struct {
	Labels map[string]string
	Env    []string
}

// Image is the abstraction of the docker client provider (formerly go-dockerclient, now either
// engine-api or kube docker client) Image type that is leveraged by s2i or origin
type Image struct {
	ID string
	*ContainerConfig
	Config *ContainerConfig
}

// DockerNetworkMode specifies the network mode setting for the docker container
type DockerNetworkMode string

const (
	// DockerNetworkModeHost places the container in the default (host) network namespace.
	DockerNetworkModeHost DockerNetworkMode = "host"
	// DockerNetworkModeBridge instructs docker to create a network namespace for this container connected to the docker0 bridge via a veth-pair.
	DockerNetworkModeBridge DockerNetworkMode = "bridge"
	// DockerNetworkModeContainerPrefix is the string prefix used by NewDockerNetworkModeContainer.
	DockerNetworkModeContainerPrefix string = "container:"
	// DockerNetworkModeNetworkNamespacePrefix is the string prefix used when sharing a namespace from a CRI-O container.
	DockerNetworkModeNetworkNamespacePrefix string = "netns:"
)

// NewDockerNetworkModeContainer creates a DockerNetworkMode value which instructs docker to place the container in the network namespace of an existing container.
// It can be used, for instance, to place the s2i container in the network namespace of the infrastructure container of a k8s pod.
func NewDockerNetworkModeContainer(id string) DockerNetworkMode {
	return DockerNetworkMode(DockerNetworkModeContainerPrefix + id)
}

// PullPolicy specifies a type for the method used to retrieve the Docker image
type PullPolicy string

// String implements the String() function of pflags.Value so this can be used as
// command line parameter.
// This method is really used just to show the default value when printing help.
// It will not default the configuration.
func (p *PullPolicy) String() string {
	if len(string(*p)) == 0 {
		return string(DefaultBuilderPullPolicy)
	}
	return string(*p)
}

// Type implements the Type() function of pflags.Value interface
func (p *PullPolicy) Type() string {
	return "string"
}

// Set implements the Set() function of pflags.Value interface
// The valid options are "always", "never" or "if-not-present"
func (p *PullPolicy) Set(v string) error {
	switch v {
	case "always":
		*p = PullAlways
	case "never":
		*p = PullNever
	case "if-not-present":
		*p = PullIfNotPresent
	default:
		return fmt.Errorf("invalid value %q, valid values are: always, never or if-not-present", v)
	}
	return nil
}

// IsInvalidFilename verifies if the provided filename contains malicious
// characters.
func IsInvalidFilename(name string) bool {
	return strings.ContainsAny(name, invalidFilenameCharacters)
}

// Set implements the Set() function of pflags.Value interface.
// This function parses the string that contains source:destination pair.
// When the destination is not specified, the source get copied into current
// working directory in container.
func (l *VolumeList) Set(value string) error {
	volumes := strings.Split(value, ";")
	newVols := make([]VolumeSpec, len(volumes))
	for i, v := range volumes {
		spec, err := l.parseSpec(v)
		if err != nil {
			return err
		}
		newVols[i] = *spec
	}
	*l = append(*l, newVols...)
	return nil
}

func (l *VolumeList) parseSpec(value string) (*VolumeSpec, error) {
	if len(value) == 0 {
		return nil, errors.New("invalid format, must be source:destination")
	}
	var mount []string
	pos := strings.LastIndex(value, ":")
	if pos == -1 {
		mount = []string{value, ""}
	} else {
		mount = []string{value[:pos], value[pos+1:]}
	}
	mount[0] = strings.Trim(mount[0], `"'`)
	mount[1] = strings.Trim(mount[1], `"'`)
	s := &VolumeSpec{Source: filepath.Clean(mount[0]), Destination: filepath.ToSlash(filepath.Clean(mount[1]))}
	if IsInvalidFilename(s.Source) || IsInvalidFilename(s.Destination) {
		return nil, fmt.Errorf("invalid characters in filename: %q", value)
	}
	return s, nil
}

// String implements the String() function of pflags.Value interface.
func (l *VolumeList) String() string {
	result := []string{}
	for _, i := range *l {
		result = append(result, strings.Join([]string{i.Source, i.Destination}, ":"))
	}
	return strings.Join(result, ",")
}

// Type implements the Type() function of pflags.Value interface.
func (l *VolumeList) Type() string {
	return "string"
}

// Set implements the Set() function of pflags.Value interface.
func (e *EnvironmentList) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 || len(parts[0]) == 0 {
		return fmt.Errorf("invalid environment format %q, must be NAME=VALUE", value)
	}
	*e = append(*e, EnvironmentSpec{
		Name:  strings.TrimSpace(parts[0]),
		Value: strings.TrimSpace(parts[1]),
	})
	return nil
}

// String implements the String() function of pflags.Value interface.
func (e *EnvironmentList) String() string {
	result := []string{}
	for _, i := range *e {
		result = append(result, strings.Join([]string{i.Name, i.Value}, "="))
	}
	return strings.Join(result, ",")
}

// Type implements the Type() function of pflags.Value interface.
func (e *EnvironmentList) Type() string {
	return "string"
}

// AsBinds converts the list of volume definitions to go-dockerclient compatible
// list of bind mounts.
func (l *VolumeList) AsBinds() []string {
	result := make([]string, len(*l))
	for index, v := range *l {
		result[index] = strings.Join([]string{v.Source, v.Destination}, ":")
	}
	return result
}
