//go:build !ignore_autogenerated

/*
Copyright  The KubeStellar Authors.

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

package v1alpha1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Binding) DeepCopyInto(out *Binding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Binding.
func (in *Binding) DeepCopy() *Binding {
	if in == nil {
		return nil
	}
	out := new(Binding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Binding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingList) DeepCopyInto(out *BindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Binding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingList.
func (in *BindingList) DeepCopy() *BindingList {
	if in == nil {
		return nil
	}
	out := new(BindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingPolicy) DeepCopyInto(out *BindingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingPolicy.
func (in *BindingPolicy) DeepCopy() *BindingPolicy {
	if in == nil {
		return nil
	}
	out := new(BindingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingPolicyCondition) DeepCopyInto(out *BindingPolicyCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingPolicyCondition.
func (in *BindingPolicyCondition) DeepCopy() *BindingPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(BindingPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingPolicyList) DeepCopyInto(out *BindingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BindingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingPolicyList.
func (in *BindingPolicyList) DeepCopy() *BindingPolicyList {
	if in == nil {
		return nil
	}
	out := new(BindingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingPolicySpec) DeepCopyInto(out *BindingPolicySpec) {
	*out = *in
	if in.ClusterSelectors != nil {
		in, out := &in.ClusterSelectors, &out.ClusterSelectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Downsync != nil {
		in, out := &in.Downsync, &out.Downsync
		*out = make([]DownsyncPolicyClause, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingPolicySpec.
func (in *BindingPolicySpec) DeepCopy() *BindingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BindingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingPolicyStatus) DeepCopyInto(out *BindingPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]BindingPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingPolicyStatus.
func (in *BindingPolicyStatus) DeepCopy() *BindingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BindingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingSpec) DeepCopyInto(out *BindingSpec) {
	*out = *in
	in.Workload.DeepCopyInto(&out.Workload)
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]Destination, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingSpec.
func (in *BindingSpec) DeepCopy() *BindingSpec {
	if in == nil {
		return nil
	}
	out := new(BindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingStatus) DeepCopyInto(out *BindingStatus) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingStatus.
func (in *BindingStatus) DeepCopy() *BindingStatus {
	if in == nil {
		return nil
	}
	out := new(BindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterScopeDownsyncClause) DeepCopyInto(out *ClusterScopeDownsyncClause) {
	*out = *in
	out.ClusterScopeDownsyncObject = in.ClusterScopeDownsyncObject
	if in.StatusCollection != nil {
		in, out := &in.StatusCollection, &out.StatusCollection
		*out = new(StatusCollection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterScopeDownsyncClause.
func (in *ClusterScopeDownsyncClause) DeepCopy() *ClusterScopeDownsyncClause {
	if in == nil {
		return nil
	}
	out := new(ClusterScopeDownsyncClause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterScopeDownsyncObject) DeepCopyInto(out *ClusterScopeDownsyncObject) {
	*out = *in
	out.GroupVersionResource = in.GroupVersionResource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterScopeDownsyncObject.
func (in *ClusterScopeDownsyncObject) DeepCopy() *ClusterScopeDownsyncObject {
	if in == nil {
		return nil
	}
	out := new(ClusterScopeDownsyncObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CombinedStatus) DeepCopyInto(out *CombinedStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]NamedStatusCombination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CombinedStatus.
func (in *CombinedStatus) DeepCopy() *CombinedStatus {
	if in == nil {
		return nil
	}
	out := new(CombinedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CombinedStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CombinedStatusList) DeepCopyInto(out *CombinedStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CombinedStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CombinedStatusList.
func (in *CombinedStatusList) DeepCopy() *CombinedStatusList {
	if in == nil {
		return nil
	}
	out := new(CombinedStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CombinedStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTransform) DeepCopyInto(out *CustomTransform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTransform.
func (in *CustomTransform) DeepCopy() *CustomTransform {
	if in == nil {
		return nil
	}
	out := new(CustomTransform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomTransform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTransformList) DeepCopyInto(out *CustomTransformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomTransform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTransformList.
func (in *CustomTransformList) DeepCopy() *CustomTransformList {
	if in == nil {
		return nil
	}
	out := new(CustomTransformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomTransformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTransformSpec) DeepCopyInto(out *CustomTransformSpec) {
	*out = *in
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTransformSpec.
func (in *CustomTransformSpec) DeepCopy() *CustomTransformSpec {
	if in == nil {
		return nil
	}
	out := new(CustomTransformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTransformStatus) DeepCopyInto(out *CustomTransformStatus) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTransformStatus.
func (in *CustomTransformStatus) DeepCopy() *CustomTransformStatus {
	if in == nil {
		return nil
	}
	out := new(CustomTransformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownsyncObjectClauses) DeepCopyInto(out *DownsyncObjectClauses) {
	*out = *in
	if in.ClusterScope != nil {
		in, out := &in.ClusterScope, &out.ClusterScope
		*out = make([]ClusterScopeDownsyncClause, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceScope != nil {
		in, out := &in.NamespaceScope, &out.NamespaceScope
		*out = make([]NamespaceScopeDownsyncClause, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownsyncObjectClauses.
func (in *DownsyncObjectClauses) DeepCopy() *DownsyncObjectClauses {
	if in == nil {
		return nil
	}
	out := new(DownsyncObjectClauses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownsyncObjectTest) DeepCopyInto(out *DownsyncObjectTest) {
	*out = *in
	if in.APIGroup != nil {
		in, out := &in.APIGroup, &out.APIGroup
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceSelectors != nil {
		in, out := &in.NamespaceSelectors, &out.NamespaceSelectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObjectSelectors != nil {
		in, out := &in.ObjectSelectors, &out.ObjectSelectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObjectNames != nil {
		in, out := &in.ObjectNames, &out.ObjectNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownsyncObjectTest.
func (in *DownsyncObjectTest) DeepCopy() *DownsyncObjectTest {
	if in == nil {
		return nil
	}
	out := new(DownsyncObjectTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownsyncPolicyClause) DeepCopyInto(out *DownsyncPolicyClause) {
	*out = *in
	in.DownsyncObjectTest.DeepCopyInto(&out.DownsyncObjectTest)
	if in.StatusCollection != nil {
		in, out := &in.StatusCollection, &out.StatusCollection
		*out = new(StatusCollection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownsyncPolicyClause.
func (in *DownsyncPolicyClause) DeepCopy() *DownsyncPolicyClause {
	if in == nil {
		return nil
	}
	out := new(DownsyncPolicyClause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressionContext) DeepCopyInto(out *ExpressionContext) {
	*out = *in
	out.Inventory = in.Inventory
	in.Obj.DeepCopyInto(&out.Obj)
	in.Returned.DeepCopyInto(&out.Returned)
	in.Propagation.DeepCopyInto(&out.Propagation)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressionContext.
func (in *ExpressionContext) DeepCopy() *ExpressionContext {
	if in == nil {
		return nil
	}
	out := new(ExpressionContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InventoryRecord) DeepCopyInto(out *InventoryRecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InventoryRecord.
func (in *InventoryRecord) DeepCopy() *InventoryRecord {
	if in == nil {
		return nil
	}
	out := new(InventoryRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedAggregator) DeepCopyInto(out *NamedAggregator) {
	*out = *in
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(Expression)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedAggregator.
func (in *NamedAggregator) DeepCopy() *NamedAggregator {
	if in == nil {
		return nil
	}
	out := new(NamedAggregator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedExpression) DeepCopyInto(out *NamedExpression) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedExpression.
func (in *NamedExpression) DeepCopy() *NamedExpression {
	if in == nil {
		return nil
	}
	out := new(NamedExpression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedStatusCombination) DeepCopyInto(out *NamedStatusCombination) {
	*out = *in
	if in.ColumnNames != nil {
		in, out := &in.ColumnNames, &out.ColumnNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Rows != nil {
		in, out := &in.Rows, &out.Rows
		*out = make([]StatusCombinationRow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedStatusCombination.
func (in *NamedStatusCombination) DeepCopy() *NamedStatusCombination {
	if in == nil {
		return nil
	}
	out := new(NamedStatusCombination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceScopeDownsyncClause) DeepCopyInto(out *NamespaceScopeDownsyncClause) {
	*out = *in
	out.NamespaceScopeDownsyncObject = in.NamespaceScopeDownsyncObject
	if in.StatusCollection != nil {
		in, out := &in.StatusCollection, &out.StatusCollection
		*out = new(StatusCollection)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceScopeDownsyncClause.
func (in *NamespaceScopeDownsyncClause) DeepCopy() *NamespaceScopeDownsyncClause {
	if in == nil {
		return nil
	}
	out := new(NamespaceScopeDownsyncClause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceScopeDownsyncObject) DeepCopyInto(out *NamespaceScopeDownsyncObject) {
	*out = *in
	out.GroupVersionResource = in.GroupVersionResource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceScopeDownsyncObject.
func (in *NamespaceScopeDownsyncObject) DeepCopy() *NamespaceScopeDownsyncObject {
	if in == nil {
		return nil
	}
	out := new(NamespaceScopeDownsyncObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropagationData) DeepCopyInto(out *PropagationData) {
	*out = *in
	in.LastReturnedUpdateTimestamp.DeepCopyInto(&out.LastReturnedUpdateTimestamp)
	in.LastGeneration = out.LastGeneration
	in.LastGenerationIsApplied = out.LastGenerationIsApplied
	in.LastCurrencyUpdateTime.DeepCopyInto(&in.LastCurrencyUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropagationData.
func (in *PropagationData) DeepCopy() *PropagationData {
	if in == nil {
		return nil
	}
	out := new(PropagationData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReturnedState) DeepCopyInto(out *ReturnedState) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReturnedState.
func (in *ReturnedState) DeepCopy() *ReturnedState {
	if in == nil {
		return nil
	}
	out := new(ReturnedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCollection) DeepCopyInto(out *StatusCollection) {
	*out = *in
	if in.StatusCollectors != nil {
		in, out := &in.StatusCollectors, &out.StatusCollectors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCollection.
func (in *StatusCollection) DeepCopy() *StatusCollection {
	if in == nil {
		return nil
	}
	out := new(StatusCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCollector) DeepCopyInto(out *StatusCollector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCollector.
func (in *StatusCollector) DeepCopy() *StatusCollector {
	if in == nil {
		return nil
	}
	out := new(StatusCollector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusCollector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCollectorList) DeepCopyInto(out *StatusCollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatusCollector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCollectorList.
func (in *StatusCollectorList) DeepCopy() *StatusCollectorList {
	if in == nil {
		return nil
	}
	out := new(StatusCollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StatusCollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCollectorSpec) DeepCopyInto(out *StatusCollectorSpec) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(Expression)
		**out = **in
	}
	if in.GroupBy != nil {
		in, out := &in.GroupBy, &out.GroupBy
		*out = make([]NamedExpression, len(*in))
		copy(*out, *in)
	}
	if in.CombinedFields != nil {
		in, out := &in.CombinedFields, &out.CombinedFields
		*out = make([]NamedAggregator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Select != nil {
		in, out := &in.Select, &out.Select
		*out = make([]NamedExpression, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCollectorSpec.
func (in *StatusCollectorSpec) DeepCopy() *StatusCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(StatusCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCollectorStatus) DeepCopyInto(out *StatusCollectorStatus) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCollectorStatus.
func (in *StatusCollectorStatus) DeepCopy() *StatusCollectorStatus {
	if in == nil {
		return nil
	}
	out := new(StatusCollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCombinationRow) DeepCopyInto(out *StatusCombinationRow) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]Value, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCombinationRow.
func (in *StatusCombinationRow) DeepCopy() *StatusCombinationRow {
	if in == nil {
		return nil
	}
	out := new(StatusCombinationRow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Value) DeepCopyInto(out *Value) {
	*out = *in
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = new(string)
		**out = **in
	}
	if in.Number != nil {
		in, out := &in.Number, &out.Number
		*out = new(string)
		**out = **in
	}
	if in.Bool != nil {
		in, out := &in.Bool, &out.Bool
		*out = new(bool)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Array != nil {
		in, out := &in.Array, &out.Array
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Value.
func (in *Value) DeepCopy() *Value {
	if in == nil {
		return nil
	}
	out := new(Value)
	in.DeepCopyInto(out)
	return out
}
