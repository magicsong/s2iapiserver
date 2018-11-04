package v1alpha1

import (
	"strings"

	"github.com/docker/distribution/reference"
	"github.com/magicsong/s2irun/pkg/api"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateConfig returns a list of error from validation.
func ValidateConfig(config *api.Config) field.ErrorList {
	allErrs := field.ErrorList{}
	if config == nil {
		allErrs = append(allErrs, field.Required(field.NewPath("spec", "config"), "Config must not be empty"))
		return allErrs
	}
	if len(config.BuilderImage) == 0 {
		allErrs = append(allErrs, field.Required(field.NewPath("spec", "config", "builderImage"), "BuilderImage must not be empty"))
	}
	if config.Labels != nil {
		for k := range config.Labels {
			if len(k) == 0 {
				allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "config", "lable"), config.Labels, "Labels value must not be empty"))
			}
		}
	}
	if config.Tag != "" {
		if err := validateDockerReference(config.Tag); err != nil {
			allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "config", "tag"), config.Tag, err.Error()))
		}
	}
	return allErrs
}

// validateDockerNetworkMode checks wether the network mode conforms to the docker remote API specification (v1.19)
// Supported values are: bridge, host, container:<name|id>, and netns:/proc/<pid>/ns/net
func validateDockerNetworkMode(mode api.DockerNetworkMode) bool {
	switch mode {
	case api.DockerNetworkModeBridge, api.DockerNetworkModeHost:
		return true
	}
	if strings.HasPrefix(string(mode), api.DockerNetworkModeContainerPrefix) {
		return true
	}
	if strings.HasPrefix(string(mode), api.DockerNetworkModeNetworkNamespacePrefix) {
		return true
	}
	return false
}

func validateDockerReference(ref string) error {
	_, err := reference.Parse(ref)
	return err
}
