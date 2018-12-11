package v1alpha1

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
