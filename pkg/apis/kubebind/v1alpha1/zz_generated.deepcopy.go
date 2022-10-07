//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubectl Bind contributors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"

	conditionsv1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/third_party/conditions/apis/conditions/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBinding) DeepCopyInto(out *ClusterBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBinding.
func (in *ClusterBinding) DeepCopy() *ClusterBinding {
	if in == nil {
		return nil
	}
	out := new(ClusterBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBindingList) DeepCopyInto(out *ClusterBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBindingList.
func (in *ClusterBindingList) DeepCopy() *ClusterBindingList {
	if in == nil {
		return nil
	}
	out := new(ClusterBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBindingSpec) DeepCopyInto(out *ClusterBindingSpec) {
	*out = *in
	in.KubeconfigSecretRef.DeepCopyInto(&out.KubeconfigSecretRef)
	in.ServiceProviderSpec.DeepCopyInto(&out.ServiceProviderSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBindingSpec.
func (in *ClusterBindingSpec) DeepCopy() *ClusterBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBindingStatus) DeepCopyInto(out *ClusterBindingStatus) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	out.HeartbeatInterval = in.HeartbeatInterval
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBindingStatus.
func (in *ClusterBindingStatus) DeepCopy() *ClusterBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCRequestSpec) DeepCopyInto(out *OIDCRequestSpec) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCRequestSpec.
func (in *OIDCRequestSpec) DeepCopy() *OIDCRequestSpec {
	if in == nil {
		return nil
	}
	out := new(OIDCRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCResponseSpec) DeepCopyInto(out *OIDCResponseSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCResponseSpec.
func (in *OIDCResponseSpec) DeepCopy() *OIDCResponseSpec {
	if in == nil {
		return nil
	}
	out := new(OIDCResponseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionClaim) DeepCopyInto(out *PermissionClaim) {
	*out = *in
	out.GroupResource = in.GroupResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionClaim.
func (in *PermissionClaim) DeepCopy() *PermissionClaim {
	if in == nil {
		return nil
	}
	out := new(PermissionClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingExportStatus) DeepCopyInto(out *ServiceBindingExportStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingExportStatus.
func (in *ServiceBindingExportStatus) DeepCopy() *ServiceBindingExportStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequest) DeepCopyInto(out *ServiceBindingRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ServiceBindingRequestSpec.DeepCopyInto(&out.ServiceBindingRequestSpec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequest.
func (in *ServiceBindingRequest) DeepCopy() *ServiceBindingRequest {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequestList) DeepCopyInto(out *ServiceBindingRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBindingRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequestList.
func (in *ServiceBindingRequestList) DeepCopy() *ServiceBindingRequestList {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequestSpec) DeepCopyInto(out *ServiceBindingRequestSpec) {
	*out = *in
	if in.OIDCRequestSpec != nil {
		in, out := &in.OIDCRequestSpec, &out.OIDCRequestSpec
		*out = new(OIDCRequestSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDCResponseSpec != nil {
		in, out := &in.OIDCResponseSpec, &out.OIDCResponseSpec
		*out = new(OIDCResponseSpec)
		**out = **in
	}
	in.ServiceProviderSpec.DeepCopyInto(&out.ServiceProviderSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequestSpec.
func (in *ServiceBindingRequestSpec) DeepCopy() *ServiceBindingRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequestStatus) DeepCopyInto(out *ServiceBindingRequestStatus) {
	*out = *in
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequestStatus.
func (in *ServiceBindingRequestStatus) DeepCopy() *ServiceBindingRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExport) DeepCopyInto(out *ServiceExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExport.
func (in *ServiceExport) DeepCopy() *ServiceExport {
	if in == nil {
		return nil
	}
	out := new(ServiceExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportList) DeepCopyInto(out *ServiceExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportList.
func (in *ServiceExportList) DeepCopy() *ServiceExportList {
	if in == nil {
		return nil
	}
	out := new(ServiceExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportSepc) DeepCopyInto(out *ServiceExportSepc) {
	*out = *in
	if in.PermissionClaims != nil {
		in, out := &in.PermissionClaims, &out.PermissionClaims
		*out = make([]PermissionClaim, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportSepc.
func (in *ServiceExportSepc) DeepCopy() *ServiceExportSepc {
	if in == nil {
		return nil
	}
	out := new(ServiceExportSepc)
	in.DeepCopyInto(out)
	return out
}
