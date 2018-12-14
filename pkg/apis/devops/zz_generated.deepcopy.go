// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package devops

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CGroupLimits) DeepCopyInto(out *CGroupLimits) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CGroupLimits.
func (in *CGroupLimits) DeepCopy() *CGroupLimits {
	if in == nil {
		return nil
	}
	out := new(CGroupLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerConfig) DeepCopyInto(out *DockerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerConfig.
func (in *DockerConfig) DeepCopy() *DockerConfig {
	if in == nil {
		return nil
	}
	out := new(DockerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Payload) DeepCopyInto(out *Payload) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Payload.
func (in *Payload) DeepCopy() *Payload {
	if in == nil {
		return nil
	}
	out := new(Payload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfig) DeepCopyInto(out *ProxyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfig.
func (in *ProxyConfig) DeepCopy() *ProxyConfig {
	if in == nil {
		return nil
	}
	out := new(ProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2IRunResult) DeepCopyInto(out *S2IRunResult) {
	*out = *in
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2IRunResult.
func (in *S2IRunResult) DeepCopy() *S2IRunResult {
	if in == nil {
		return nil
	}
	out := new(S2IRunResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilder) DeepCopyInto(out *S2iBuilder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilder.
func (in *S2iBuilder) DeepCopy() *S2iBuilder {
	if in == nil {
		return nil
	}
	out := new(S2iBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iBuilder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderList) DeepCopyInto(out *S2iBuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S2iBuilder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderList.
func (in *S2iBuilderList) DeepCopy() *S2iBuilderList {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iBuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderSpec) DeepCopyInto(out *S2iBuilderSpec) {
	*out = *in
	out.RuntimeAuthentication = in.RuntimeAuthentication
	if in.RuntimeArtifacts != nil {
		in, out := &in.RuntimeArtifacts, &out.RuntimeArtifacts
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.DockerConfig != nil {
		in, out := &in.DockerConfig, &out.DockerConfig
		*out = new(DockerConfig)
		**out = **in
	}
	out.PullAuthentication = in.PullAuthentication
	out.PushAuthentication = in.PushAuthentication
	out.IncrementalAuthentication = in.IncrementalAuthentication
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]EnvironmentSpec, len(*in))
		copy(*out, *in)
	}
	if in.Injections != nil {
		in, out := &in.Injections, &out.Injections
		*out = make([]VolumeSpec, len(*in))
		copy(*out, *in)
	}
	if in.CGroupLimits != nil {
		in, out := &in.CGroupLimits, &out.CGroupLimits
		*out = new(CGroupLimits)
		**out = **in
	}
	if in.DropCapabilities != nil {
		in, out := &in.DropCapabilities, &out.DropCapabilities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScriptDownloadProxyConfig != nil {
		in, out := &in.ScriptDownloadProxyConfig, &out.ScriptDownloadProxyConfig
		*out = new(ProxyConfig)
		**out = **in
	}
	if in.BuildVolumes != nil {
		in, out := &in.BuildVolumes, &out.BuildVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityOpt != nil {
		in, out := &in.SecurityOpt, &out.SecurityOpt
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AddHost != nil {
		in, out := &in.AddHost, &out.AddHost
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderSpec.
func (in *S2iBuilderSpec) DeepCopy() *S2iBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iBuilderStatus) DeepCopyInto(out *S2iBuilderStatus) {
	*out = *in
	if in.LastRunName != nil {
		in, out := &in.LastRunName, &out.LastRunName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iBuilderStatus.
func (in *S2iBuilderStatus) DeepCopy() *S2iBuilderStatus {
	if in == nil {
		return nil
	}
	out := new(S2iBuilderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRun) DeepCopyInto(out *S2iRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRun.
func (in *S2iRun) DeepCopy() *S2iRun {
	if in == nil {
		return nil
	}
	out := new(S2iRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRunList) DeepCopyInto(out *S2iRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S2iRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRunList.
func (in *S2iRunList) DeepCopy() *S2iRunList {
	if in == nil {
		return nil
	}
	out := new(S2iRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S2iRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRunSpec) DeepCopyInto(out *S2iRunSpec) {
	*out = *in
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]EnvironmentSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRunSpec.
func (in *S2iRunSpec) DeepCopy() *S2iRunSpec {
	if in == nil {
		return nil
	}
	out := new(S2iRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2iRunStatus) DeepCopyInto(out *S2iRunStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(Trigger)
		(*in).DeepCopyInto(*out)
	}
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		*out = new(S2IRunResult)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2iRunStatus.
func (in *S2iRunStatus) DeepCopy() *S2iRunStatus {
	if in == nil {
		return nil
	}
	out := new(S2iRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.Payload != nil {
		in, out := &in.Payload, &out.Payload
		*out = new(Payload)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
