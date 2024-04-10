//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 The Volcano Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	batchv1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.CreateTimestamp.DeepCopyInto(&out.CreateTimestamp)
	if in.RunningDuration != nil {
		in, out := &in.RunningDuration, &out.RunningDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TaskStatusCount != nil {
		in, out := &in.TaskStatusCount, &out.TaskStatusCount
		*out = make(map[string]batchv1alpha1.TaskState, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependsOn) DeepCopyInto(out *DependsOn) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(Probe)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependsOn.
func (in *DependsOn) DeepCopy() *DependsOn {
	if in == nil {
		return nil
	}
	out := new(DependsOn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flow) DeepCopyInto(out *Flow) {
	*out = *in
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = new(DependsOn)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flow.
func (in *Flow) DeepCopy() *Flow {
	if in == nil {
		return nil
	}
	out := new(Flow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpGet) DeepCopyInto(out *HttpGet) {
	*out = *in
	out.HTTPHeader = in.HTTPHeader
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpGet.
func (in *HttpGet) DeepCopy() *HttpGet {
	if in == nil {
		return nil
	}
	out := new(HttpGet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobFlow) DeepCopyInto(out *JobFlow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobFlow.
func (in *JobFlow) DeepCopy() *JobFlow {
	if in == nil {
		return nil
	}
	out := new(JobFlow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobFlow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobFlowList) DeepCopyInto(out *JobFlowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JobFlow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobFlowList.
func (in *JobFlowList) DeepCopy() *JobFlowList {
	if in == nil {
		return nil
	}
	out := new(JobFlowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobFlowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobFlowSpec) DeepCopyInto(out *JobFlowSpec) {
	*out = *in
	if in.Flows != nil {
		in, out := &in.Flows, &out.Flows
		*out = make([]Flow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobFlowSpec.
func (in *JobFlowSpec) DeepCopy() *JobFlowSpec {
	if in == nil {
		return nil
	}
	out := new(JobFlowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobFlowStatus) DeepCopyInto(out *JobFlowStatus) {
	*out = *in
	if in.PendingJobs != nil {
		in, out := &in.PendingJobs, &out.PendingJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RunningJobs != nil {
		in, out := &in.RunningJobs, &out.RunningJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FailedJobs != nil {
		in, out := &in.FailedJobs, &out.FailedJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CompletedJobs != nil {
		in, out := &in.CompletedJobs, &out.CompletedJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TerminatedJobs != nil {
		in, out := &in.TerminatedJobs, &out.TerminatedJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UnKnowJobs != nil {
		in, out := &in.UnKnowJobs, &out.UnKnowJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.JobStatusList != nil {
		in, out := &in.JobStatusList, &out.JobStatusList
		*out = make([]JobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]Condition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.State = in.State
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobFlowStatus.
func (in *JobFlowStatus) DeepCopy() *JobFlowStatus {
	if in == nil {
		return nil
	}
	out := new(JobFlowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobRunningHistory) DeepCopyInto(out *JobRunningHistory) {
	*out = *in
	in.StartTimestamp.DeepCopyInto(&out.StartTimestamp)
	in.EndTimestamp.DeepCopyInto(&out.EndTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobRunningHistory.
func (in *JobRunningHistory) DeepCopy() *JobRunningHistory {
	if in == nil {
		return nil
	}
	out := new(JobRunningHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatus) DeepCopyInto(out *JobStatus) {
	*out = *in
	in.StartTimestamp.DeepCopyInto(&out.StartTimestamp)
	in.EndTimestamp.DeepCopyInto(&out.EndTimestamp)
	if in.RunningHistories != nil {
		in, out := &in.RunningHistories, &out.RunningHistories
		*out = make([]JobRunningHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStatus.
func (in *JobStatus) DeepCopy() *JobStatus {
	if in == nil {
		return nil
	}
	out := new(JobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplate) DeepCopyInto(out *JobTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplate.
func (in *JobTemplate) DeepCopy() *JobTemplate {
	if in == nil {
		return nil
	}
	out := new(JobTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplateList) DeepCopyInto(out *JobTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JobTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplateList.
func (in *JobTemplateList) DeepCopy() *JobTemplateList {
	if in == nil {
		return nil
	}
	out := new(JobTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplateSpec) DeepCopyInto(out *JobTemplateSpec) {
	*out = *in
	in.JobSpec.DeepCopyInto(&out.JobSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplateSpec.
func (in *JobTemplateSpec) DeepCopy() *JobTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(JobTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplateStatus) DeepCopyInto(out *JobTemplateStatus) {
	*out = *in
	if in.JobDependsOnList != nil {
		in, out := &in.JobDependsOnList, &out.JobDependsOnList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplateStatus.
func (in *JobTemplateStatus) DeepCopy() *JobTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(JobTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
	if in.HttpGetList != nil {
		in, out := &in.HttpGetList, &out.HttpGetList
		*out = make([]HttpGet, len(*in))
		copy(*out, *in)
	}
	if in.TcpSocketList != nil {
		in, out := &in.TcpSocketList, &out.TcpSocketList
		*out = make([]TcpSocket, len(*in))
		copy(*out, *in)
	}
	if in.TaskStatusList != nil {
		in, out := &in.TaskStatusList, &out.TaskStatusList
		*out = make([]TaskStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *State) DeepCopyInto(out *State) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new State.
func (in *State) DeepCopy() *State {
	if in == nil {
		return nil
	}
	out := new(State)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcpSocket) DeepCopyInto(out *TcpSocket) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcpSocket.
func (in *TcpSocket) DeepCopy() *TcpSocket {
	if in == nil {
		return nil
	}
	out := new(TcpSocket)
	in.DeepCopyInto(out)
	return out
}
