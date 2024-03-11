/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JobSetSpec defines the desired state of JobSet
type JobSetSpec struct {
	// ReplicatedJobs is a group of volcano jobs managed by JobSet.
	// +listType=map
	// +listMapKey=name
	ReplicatedJobs []ReplicatedJob `json:"replicatedJobs,omitempty"`

	// The minimal available jobs to run for this JobSet.
	// The default is the number of all jobs.
	// +optional
	MinAvailable int32 `json:"minAvailable,omitempty" protobuf:"bytes,2,opt,name=minAvailable"`

	//Specifies the queue that will be used in the scheduler, "default" queue is used this leaves empty.
	// +optional
	Queue string `json:"queue,omitempty" protobuf:"bytes,8,opt,name=queue"`

	// If specified, indicates the job's priority.
	// +optional
	PriorityClassName string `json:"priorityClassName,omitempty" protobuf:"bytes,11,opt,name=priorityClassName"`

	// SuccessPolicy configures when to declare the JobSet as
	// succeeded.
	// The JobSet is always declared succeeded if all jobs in the set
	// finished with status complete.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	SuccessPolicy *SuccessPolicy `json:"successPolicy,omitempty"`

	// FailurePolicy, if set, configures when to declare the JobSet as
	// failed.
	// The JobSet is always declared failed if any job in the set
	// finished with status failed.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	FailurePolicy *FailurePolicy `json:"failurePolicy,omitempty"`

	// StartupPolicy, if set, configures in what order jobs must be started
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	StartupPolicy *StartupPolicy `json:"startupPolicy,omitempty"`

	// Suspend suspends all running child Jobs when set to true.
	Suspend *bool `json:"suspend,omitempty"`
}

// JobSetStatus defines the observed state of JobSet
type JobSetStatus struct {
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Restarts tracks the number of times the JobSet has restarted (i.e. recreated in case of RecreateAll policy).
	Restarts int32 `json:"restarts,omitempty"`

	// ReplicatedJobsStatus track the number of JobsReady for each replicatedJob.
	// +optional
	// +listType=map
	// +listMapKey=name
	ReplicatedJobsStatus []ReplicatedJobStatus `json:"replicatedJobsStatus,omitempty"`
}

// ReplicatedJobStatus defines the observed ReplicatedJobs Readiness.
type ReplicatedJobStatus struct {
	Name      string `json:"name"`
	Ready     int32  `json:"ready"`
	Succeeded int32  `json:"succeeded"`
	Failed    int32  `json:"failed"`
	Active    int32  `json:"active"`
	Suspended int32  `json:"suspended"`
}

//+genclient
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:resource:path=jobsets,shortName=jobset;js
//+kubebuilder:subresource:status

// JobSet is the Schema for the JobSets API
type JobSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JobSetSpec   `json:"spec,omitempty"`
	Status JobSetStatus `json:"status,omitempty"`
}

//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// JobSetList contains a list of JobSet
type JobSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobSet `json:"items"`
}

type ReplicatedJob struct {
	// Name is the name of the entry and will be used as a suffix for the Job name.
	Name string `json:"name"`
	// Template defines the template of the Job that will be created.
	Template JobSpec `json:"template"`

	// Replicas is the number of jobs that will be created from this ReplicatedJob's template.
	// Jobs names will be in the format: <JobSet.name>-<spec.replicatedJob.name>-<job-index>
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`
}

type JobSetConditionType string

// These are built-in conditions of a JobSet.
const (
	// JobSetComplete means the job has completed its execution.
	JobSetCompleted JobSetConditionType = "Completed"
	// JobSetFailed means the job has failed its execution.
	JobSetFailed JobSetConditionType = "Failed"
	// JobSetSuspended means the job is suspended
	JobSetSuspended JobSetConditionType = "Suspended"
	// JobSetStartupPolicyCompleted means the StartupPolicy was complete
	JobSetStartupPolicyCompleted JobSetConditionType = "StartupPolicyCompleted"
)

// Operator defines the target of a SuccessPolicy or FailurePolicy.
type Operator string

const (
	// OperatorAll applies to all jobs matching the jobSelector.
	OperatorAll Operator = "All"

	// OperatorAny applies to any single job matching the jobSelector.
	OperatorAny Operator = "Any"
)

type FailurePolicy struct {
	// MaxRestarts defines the limit on the number of JobSet restarts.
	// A restart is achieved by recreating all active child jobs.
	MaxRestarts int32 `json:"maxRestarts,omitempty"`
}

type SuccessPolicy struct {
	// Operator determines either All or Any of the selected jobs should succeed to consider the JobSet successful
	// +kubebuilder:validation:Enum=All;Any
	Operator Operator `json:"operator"`

	// TargetReplicatedJobs are the names of the replicated jobs the operator will apply to.
	// A null or empty list will apply to all replicatedJobs.
	// +optional
	// +listType=atomic
	TargetReplicatedJobs []string `json:"targetReplicatedJobs,omitempty"`
}

type StartupPolicyOptions string

const (
	// This is the default setting
	// AnyOrder means that we will start the replicated jobs
	// without any specific order.
	AnyOrder StartupPolicyOptions = "AnyOrder"
	// InOrder starts the replicated jobs in order
	// that they are listed.
	InOrder StartupPolicyOptions = "InOrder"
)

type StartupPolicy struct {
	// StartupPolicyOrder determines the startup order of the ReplicatedJobs.
	// AnyOrder means to start replicated jobs in any order.
	// InOrder means to start them as they are listed in the JobSet. A ReplicatedJob is started only
	// when all the jobs of the previous one are ready.
	// +kubebuilder:validation:Enum=AnyOrder;InOrder
	StartupPolicyOrder StartupPolicyOptions `json:"startupPolicyOrder"`
}
