//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Volcano Authors.

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
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	busv1alpha1 "volcano.sh/apis/pkg/apis/bus/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependsOn) DeepCopyInto(out *DependsOn) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = make([]string, len(*in))
		copy(*out, *in)
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
func (in *FailurePolicy) DeepCopyInto(out *FailurePolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailurePolicy.
func (in *FailurePolicy) DeepCopy() *FailurePolicy {
	if in == nil {
		return nil
	}
	out := new(FailurePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperJob) DeepCopyInto(out *HyperJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperJob.
func (in *HyperJob) DeepCopy() *HyperJob {
	if in == nil {
		return nil
	}
	out := new(HyperJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperJobList) DeepCopyInto(out *HyperJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HyperJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperJobList.
func (in *HyperJobList) DeepCopy() *HyperJobList {
	if in == nil {
		return nil
	}
	out := new(HyperJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HyperJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperJobSpec) DeepCopyInto(out *HyperJobSpec) {
	*out = *in
	if in.ReplicatedJobs != nil {
		in, out := &in.ReplicatedJobs, &out.ReplicatedJobs
		*out = make([]ReplicatedJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuccessPolicy != nil {
		in, out := &in.SuccessPolicy, &out.SuccessPolicy
		*out = new(SuccessPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(FailurePolicy)
		**out = **in
	}
	if in.StartupPolicy != nil {
		in, out := &in.StartupPolicy, &out.StartupPolicy
		*out = new(StartupPolicy)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperJobSpec.
func (in *HyperJobSpec) DeepCopy() *HyperJobSpec {
	if in == nil {
		return nil
	}
	out := new(HyperJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HyperJobStatus) DeepCopyInto(out *HyperJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReplicatedJobsStatus != nil {
		in, out := &in.ReplicatedJobsStatus, &out.ReplicatedJobsStatus
		*out = make([]ReplicatedJobStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HyperJobStatus.
func (in *HyperJobStatus) DeepCopy() *HyperJobStatus {
	if in == nil {
		return nil
	}
	out := new(HyperJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Job) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobCondition) DeepCopyInto(out *JobCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobCondition.
func (in *JobCondition) DeepCopy() *JobCondition {
	if in == nil {
		return nil
	}
	out := new(JobCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobList) DeepCopyInto(out *JobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Job, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobList.
func (in *JobList) DeepCopy() *JobList {
	if in == nil {
		return nil
	}
	out := new(JobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSpec) DeepCopyInto(out *JobSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VolumeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]TaskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]LifecyclePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.RunningEstimate != nil {
		in, out := &in.RunningEstimate, &out.RunningEstimate
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.MinSuccess != nil {
		in, out := &in.MinSuccess, &out.MinSuccess
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSpec.
func (in *JobSpec) DeepCopy() *JobSpec {
	if in == nil {
		return nil
	}
	out := new(JobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobState) DeepCopyInto(out *JobState) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobState.
func (in *JobState) DeepCopy() *JobState {
	if in == nil {
		return nil
	}
	out := new(JobState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatus) DeepCopyInto(out *JobStatus) {
	*out = *in
	in.State.DeepCopyInto(&out.State)
	if in.TaskStatusCount != nil {
		in, out := &in.TaskStatusCount, &out.TaskStatusCount
		*out = make(map[string]TaskState, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.RunningDuration != nil {
		in, out := &in.RunningDuration, &out.RunningDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ControlledResources != nil {
		in, out := &in.ControlledResources, &out.ControlledResources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JobCondition, len(*in))
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
func (in *LifecyclePolicy) DeepCopyInto(out *LifecyclePolicy) {
	*out = *in
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]busv1alpha1.Event, len(*in))
		copy(*out, *in)
	}
	if in.ExitCode != nil {
		in, out := &in.ExitCode, &out.ExitCode
		*out = new(int32)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecyclePolicy.
func (in *LifecyclePolicy) DeepCopy() *LifecyclePolicy {
	if in == nil {
		return nil
	}
	out := new(LifecyclePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedJob) DeepCopyInto(out *ReplicatedJob) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedJob.
func (in *ReplicatedJob) DeepCopy() *ReplicatedJob {
	if in == nil {
		return nil
	}
	out := new(ReplicatedJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicatedJobStatus) DeepCopyInto(out *ReplicatedJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicatedJobStatus.
func (in *ReplicatedJobStatus) DeepCopy() *ReplicatedJobStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicatedJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartupPolicy) DeepCopyInto(out *StartupPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartupPolicy.
func (in *StartupPolicy) DeepCopy() *StartupPolicy {
	if in == nil {
		return nil
	}
	out := new(StartupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuccessPolicy) DeepCopyInto(out *SuccessPolicy) {
	*out = *in
	if in.TargetReplicatedJobs != nil {
		in, out := &in.TargetReplicatedJobs, &out.TargetReplicatedJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuccessPolicy.
func (in *SuccessPolicy) DeepCopy() *SuccessPolicy {
	if in == nil {
		return nil
	}
	out := new(SuccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(int32)
		**out = **in
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]LifecyclePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = new(DependsOn)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskState) DeepCopyInto(out *TaskState) {
	*out = *in
	if in.Phase != nil {
		in, out := &in.Phase, &out.Phase
		*out = make(map[corev1.PodPhase]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskState.
func (in *TaskState) DeepCopy() *TaskState {
	if in == nil {
		return nil
	}
	out := new(TaskState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	if in.VolumeClaim != nil {
		in, out := &in.VolumeClaim, &out.VolumeClaim
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
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
