//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v2beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPrepper) DeepCopyInto(out *DataPrepper) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPrepper.
func (in *DataPrepper) DeepCopy() *DataPrepper {
	if in == nil {
		return nil
	}
	out := new(DataPrepper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataPrepper) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPrepperList) DeepCopyInto(out *DataPrepperList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataPrepper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPrepperList.
func (in *DataPrepperList) DeepCopy() *DataPrepperList {
	if in == nil {
		return nil
	}
	out := new(DataPrepperList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataPrepperList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPrepperSpec) DeepCopyInto(out *DataPrepperSpec) {
	*out = *in
	if in.ImageSpec != nil {
		in, out := &in.ImageSpec, &out.ImageSpec
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultRepo != nil {
		in, out := &in.DefaultRepo, &out.DefaultRepo
		*out = new(string)
		**out = **in
	}
	if in.Opensearch != nil {
		in, out := &in.Opensearch, &out.Opensearch
		*out = new(OpensearchSpec)
		**out = **in
	}
	if in.PasswordFrom != nil {
		in, out := &in.PasswordFrom, &out.PasswordFrom
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPrepperSpec.
func (in *DataPrepperSpec) DeepCopy() *DataPrepperSpec {
	if in == nil {
		return nil
	}
	out := new(DataPrepperSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPrepperStatus) DeepCopyInto(out *DataPrepperStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPrepperStatus.
func (in *DataPrepperStatus) DeepCopy() *DataPrepperStatus {
	if in == nil {
		return nil
	}
	out := new(DataPrepperStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageResolver) DeepCopyInto(out *ImageResolver) {
	*out = *in
	if in.DefaultRepoOverride != nil {
		in, out := &in.DefaultRepoOverride, &out.DefaultRepoOverride
		*out = new(string)
		**out = **in
	}
	if in.ImageOverride != nil {
		in, out := &in.ImageOverride, &out.ImageOverride
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageResolver.
func (in *ImageResolver) DeepCopy() *ImageResolver {
	if in == nil {
		return nil
	}
	out := new(ImageResolver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingCluster) DeepCopyInto(out *LoggingCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingCluster.
func (in *LoggingCluster) DeepCopy() *LoggingCluster {
	if in == nil {
		return nil
	}
	out := new(LoggingCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingClusterList) DeepCopyInto(out *LoggingClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoggingCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingClusterList.
func (in *LoggingClusterList) DeepCopy() *LoggingClusterList {
	if in == nil {
		return nil
	}
	out := new(LoggingClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingClusterSpec) DeepCopyInto(out *LoggingClusterSpec) {
	*out = *in
	if in.OpensearchClusterRef != nil {
		in, out := &in.OpensearchClusterRef, &out.OpensearchClusterRef
		*out = new(OpensearchClusterRef)
		**out = **in
	}
	if in.IndexUserSecret != nil {
		in, out := &in.IndexUserSecret, &out.IndexUserSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingClusterSpec.
func (in *LoggingClusterSpec) DeepCopy() *LoggingClusterSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingClusterStatus) DeepCopyInto(out *LoggingClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingClusterStatus.
func (in *LoggingClusterStatus) DeepCopy() *LoggingClusterStatus {
	if in == nil {
		return nil
	}
	out := new(LoggingClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MulticlusterRoleBinding) DeepCopyInto(out *MulticlusterRoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MulticlusterRoleBinding.
func (in *MulticlusterRoleBinding) DeepCopy() *MulticlusterRoleBinding {
	if in == nil {
		return nil
	}
	out := new(MulticlusterRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MulticlusterRoleBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MulticlusterRoleBindingList) DeepCopyInto(out *MulticlusterRoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MulticlusterRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MulticlusterRoleBindingList.
func (in *MulticlusterRoleBindingList) DeepCopy() *MulticlusterRoleBindingList {
	if in == nil {
		return nil
	}
	out := new(MulticlusterRoleBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MulticlusterRoleBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MulticlusterRoleBindingSpec) DeepCopyInto(out *MulticlusterRoleBindingSpec) {
	*out = *in
	if in.OpensearchCluster != nil {
		in, out := &in.OpensearchCluster, &out.OpensearchCluster
		*out = new(OpensearchClusterRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MulticlusterRoleBindingSpec.
func (in *MulticlusterRoleBindingSpec) DeepCopy() *MulticlusterRoleBindingSpec {
	if in == nil {
		return nil
	}
	out := new(MulticlusterRoleBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MulticlusterRoleBindingStatus) DeepCopyInto(out *MulticlusterRoleBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MulticlusterRoleBindingStatus.
func (in *MulticlusterRoleBindingStatus) DeepCopy() *MulticlusterRoleBindingStatus {
	if in == nil {
		return nil
	}
	out := new(MulticlusterRoleBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpensearchClusterRef) DeepCopyInto(out *OpensearchClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpensearchClusterRef.
func (in *OpensearchClusterRef) DeepCopy() *OpensearchClusterRef {
	if in == nil {
		return nil
	}
	out := new(OpensearchClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpensearchSpec) DeepCopyInto(out *OpensearchSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpensearchSpec.
func (in *OpensearchSpec) DeepCopy() *OpensearchSpec {
	if in == nil {
		return nil
	}
	out := new(OpensearchSpec)
	in.DeepCopyInto(out)
	return out
}
