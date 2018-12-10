package v1alpha1

import (
	"github.com/magicsong/s2iapiserver/pkg/apis/devops/constants"
)

// Config contains essential fields for performing build.
type Config struct {
	// DisplayName is a result image display-name label. This defaults to the
	// output image name.
	DisplayName string `json:"display_name,omitempty"`

	// Description is a result image description label. The default is no
	// description.
	Description string `json:"description,omitempty"`

	// BuilderImage describes which image is used for building the result images.
	BuilderImage string `json:"builder_image,omitempty"`

	// BuilderImageVersion provides optional version information about the builder image.
	BuilderImageVersion string `json:"builder_image_version,omitempty"`

	// BuilderBaseImageVersion provides optional version information about the builder base image.
	BuilderBaseImageVersion string `json:"builder_base_image_version,omitempty"`

	// RuntimeImage specifies the image that will be a base for resulting image
	// and will be used for running an application. By default, BuilderImage is
	// used for building and running, but the latter may be overridden.
	RuntimeImage string `json:"runtime_image,omitempty"`

	// RuntimeImagePullPolicy specifies when to pull a runtime image.
	RuntimeImagePullPolicy constants.PullPolicy `json:"runtime_image_pull_policy,omitempty"`

	// RuntimeAuthentication holds the authentication information for pulling the
	// runtime Docker images from private repositories.
	RuntimeAuthentication AuthConfig `json:"runtime_authentication,omitempty"`

	// RuntimeArtifacts specifies a list of source/destination pairs that will
	// be copied from builder to a runtime image. Source can be a file or
	// directory. Destination must be a directory. Regardless whether it
	// is an absolute or relative path, it will be placed into image's WORKDIR.
	// Destination also can be empty or equals to ".", in this case it just
	// refers to a root of WORKDIR.
	// In case it's empty, S2I will try to get this list from
	// io.openshift.s2i.assemble-input-files label on a RuntimeImage.
	RuntimeArtifacts []VolumeSpec `json:"runtime_artifacts,omitempty"`

	// DockerConfig describes how to access host docker daemon.
	DockerConfig *DockerConfig `json:"docker_config,omitempty"`

	// PullAuthentication holds the authentication information for pulling the
	// Docker images from private repositories
	PullAuthentication AuthConfig `json:"pull_authentication,omitempty"`

	// PullAuthentication holds the authentication information for pulling the
	// Docker images from private repositories
	PushAuthentication AuthConfig `json:"push_authentication,omitempty"`

	// IncrementalAuthentication holds the authentication information for pulling the
	// previous image from private repositories
	IncrementalAuthentication AuthConfig `json:"incremental_authentication,omitempty"`

	// DockerNetworkMode is used to set the docker network setting to --net=container:<id>
	// when the builder is invoked from a container.
	DockerNetworkMode constants.DockerNetworkMode `json:"docker_network_mode,omitempty"`

	// PreserveWorkingDir describes if working directory should be left after processing.
	PreserveWorkingDir bool `json:"preserve_working_dir,omitempty"`

	// IgnoreSubmodules determines whether we will attempt to pull in submodules
	// (via --recursive or submodule init)
	IgnoreSubmodules bool `json:"ignore_submodules,omitempty"`

	// Tag is a result image tag name.
	Tag string `json:"tag,omitempty"`

	// BuilderPullPolicy specifies when to pull the builder image
	BuilderPullPolicy constants.PullPolicy `json:"builder_pull_policy,omitempty"`

	// PreviousImagePullPolicy specifies when to pull the previously build image
	// when doing incremental build
	PreviousImagePullPolicy constants.PullPolicy `json:"previous_image_pull_policy,omitempty"`

	// Incremental describes whether to try to perform incremental build.
	Incremental bool `json:"incremental,omitempty"`

	// IncrementalFromTag sets an alternative image tag to look for existing
	// artifacts. Tag is used by default if this is not set.
	IncrementalFromTag string `json:"incremental_from_tag,omitempty"`

	// RemovePreviousImage describes if previous image should be removed after successful build.
	// This applies only to incremental builds.
	RemovePreviousImage bool `json:"remove_previous_image,omitempty"`

	// Environment is a map of environment variables to be passed to the image.
	Environment []EnvironmentSpec `json:"environment,omitempty"`

	// LabelNamespace provides the namespace under which the labels will be generated.
	LabelNamespace string `json:"label_namespace,omitempty"`

	// CallbackURL is a URL which is called upon successful build to inform about that fact.
	CallbackURL string `json:"callback_url,omitempty"`

	// ScriptsURL is a URL describing where to fetch the S2I scripts from during build process.
	// This url can be a reference within the builder image if the scheme is specified as image://
	ScriptsURL string `json:"scripts_url,omitempty"`

	// Destination specifies a location where the untar operation will place its artifacts.
	Destination string `json:"destination,omitempty"`

	// WorkingDir describes temporary directory used for downloading sources, scripts and tar operations.
	WorkingDir string `json:"working_dir,omitempty"`

	// WorkingSourceDir describes the subdirectory off of WorkingDir set up during the repo download
	// that is later used as the root for ignore processing
	WorkingSourceDir string `json:"working_source_dir,omitempty"`

	// LayeredBuild describes if this is build which layered scripts and sources on top of BuilderImage.
	LayeredBuild bool `json:"layered_build,omitempty"`

	// Operate quietly. Progress and assemble script output are not reported, only fatal errors.
	// (default: false).
	Quiet bool `json:"quiet,omitempty"`

	// ForceCopy results in only the file SCM plugin being used (i.e. no `git clone`); allows for empty directories to be included
	// in resulting image (since git does not support that).
	// (default: false).
	ForceCopy bool `json:"force_copy,omitempty"`

	// Specify a relative directory inside the application repository that should
	// be used as a root directory for the application.
	ContextDir string `json:"context_dir,omitempty"`

	// AssembleUser specifies the user to run the assemble script in container
	AssembleUser string `json:"assemble_user,omitempty"`

	// RunImage will trigger a "docker run ..." invocation of the produced image so the user
	// can see if it operates as he would expect
	RunImage bool `json:"run_image,omitempty"`

	// Usage allows for properly shortcircuiting s2i logic when `s2i usage` is invoked
	Usage bool `json:"usage,omitempty"`

	// Injections specifies a list source/destination folders that are injected to
	// the container that runs assemble.
	// All files we inject will be truncated after the assemble script finishes.
	Injections []VolumeSpec `json:"injections,omitempty"`

	// CGroupLimits describes the cgroups limits that will be applied to any containers
	// run by s2i.
	CGroupLimits *CGroupLimits `json:"cgroup_limits,omitempty"`

	// DropCapabilities contains a list of capabilities to drop when executing containers
	DropCapabilities []string `json:"drop_capabilities,omitempty"`

	// ScriptDownloadProxyConfig optionally specifies the http and https proxy
	// to use when downloading scripts
	ScriptDownloadProxyConfig *ProxyConfig `json:"script_download_proxy_config,omitempty"`

	// ExcludeRegExp contains a string representation of the regular expression desired for
	// deciding which files to exclude from the tar stream
	ExcludeRegExp string `json:"exclude_reg_exp,omitempty"`

	// BlockOnBuild prevents s2i from performing a docker build operation
	// if one is necessary to execute ONBUILD commands, or to layer source code into
	// the container for images that don't have a tar binary available, if the
	// image contains ONBUILD commands that would be executed.
	BlockOnBuild bool `json:"block_on_build,omitempty"`

	// HasOnBuild will be set to true if the builder image contains ONBUILD instructions
	HasOnBuild bool `json:"has_on_build,omitempty"`

	// BuildVolumes specifies a list of volumes to mount to container running the
	// build.
	BuildVolumes []string `json:"build_volumes,omitempty"`

	// Labels specify labels and their values to be applied to the resulting image. Label keys
	// must have non-zero length. The labels defined here override generated labels in case
	// they have the same name.
	Labels map[string]string `json:"labels,omitempty"`

	// SecurityOpt are passed as options to the docker containers launched by s2i.
	SecurityOpt []string `json:"security_opt,omitempty"`

	// KeepSymlinks indicates to copy symlinks as symlinks. Default behavior is to follow
	// symlinks and copy files by content.
	KeepSymlinks bool `json:"keep_symlinks,omitempty"`

	// AsDockerfile indicates the path where the Dockerfile should be written instead of building
	// a new image.
	AsDockerfile string `json:"as_dockerfile,omitempty"`

	// ImageWorkDir is the default working directory for the builder image.
	ImageWorkDir string `json:"image_work_dir,omitempty"`

	// ImageScriptsURL is the default location to find the assemble/run scripts for a builder image.
	// This url can be a reference within the builder image if the scheme is specified as image://
	ImageScriptsURL string `json:"image_scripts_url,omitempty"`

	// AddHost Add a line to /etc/hosts for test purpose or private use in LAN. Its format is host:IP,muliple hosts can be added  by using multiple --add-host
	AddHost []string `json:"add_host,omitempty"`

	//Export Push the result image to specify image registry in tag
	Export bool `json:"export,omitempty"`

	SourceURL string `json:"source_url,omitempty"`
}

// EnvironmentSpec specifies a single environment variable.
type EnvironmentSpec struct {
	Name  string
	Value string
}

// ProxyConfig holds proxy configuration.
type ProxyConfig struct {
	HTTPProxy  string
	HTTPSProxy string
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
