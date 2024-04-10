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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HyperJobSpec defines the desired state of HyperJob
type HyperJobSpec struct {
	// ReplicatedJobs is a group of volcano jobs managed by HyperJob.
	// +listType=map
	// +listMapKey=name
	ReplicatedJobs []ReplicatedJob `json:"replicatedJobs,omitempty"`

	// The minimal available jobs to run for this HyperJob.
	// The default is the number of all jobs.
	// +optional
	MinAvailable int32 `json:"minAvailable,omitempty" protobuf:"bytes,2,opt,name=minAvailable"`

	// SuccessPolicy configures when to declare the HyperJob as
	// succeeded.
	// The HyperJob is always declared succeeded if all jobs in the set
	// finished with status complete.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	SuccessPolicy *SuccessPolicy `json:"successPolicy,omitempty"`

	// FailurePolicy, if set, configures when to declare the HyperJob as
	// failed.
	// The HyperJob is always declared failed if any job in the set
	// finished with status failed.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	FailurePolicy *FailurePolicy `json:"failurePolicy,omitempty"`

	// StartupPolicy, if set, configures in what order jobs must be started
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	StartupPolicy *StartupPolicy `json:"startupPolicy,omitempty"`
}

// HyperJobStatus defines the observed state of HyperJob
type HyperJobStatus struct {
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Restarts tracks the number of times the HyperJob has restarted (i.e. recreated in case of RecreateAll policy).
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
}

//+genclient
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:resource:path=hyperjobs,shortName=hyperjob;jobset;js
//+kubebuilder:subresource:status

// HyperJob is the Schema for the HyperJobs API
type HyperJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HyperJobSpec   `json:"spec,omitempty"`
	Status HyperJobStatus `json:"status,omitempty"`
}

//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// HyperJobList contains a list of HyperJob
type HyperJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HyperJob `json:"items"`
}

type ReplicatedJob struct {
	// Name is the name of the entry and will be used as a suffix for the Job name.
	Name string `json:"name"`
	// Template defines the template of the Job that will be created.
	Template JobSpec `json:"template"`

	// Replicas is the number of jobs that will be created from this ReplicatedJob's template.
	// Jobs names will be in the format: <HyperJob.name>-<spec.replicatedJob.name>-<job-index>
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`
}

type HyperJobConditionType string

// These are built-in conditions of a HyperJob.
const (
	// HyperJobComplete means the job has completed its execution.
	HyperJobCompleted HyperJobConditionType = "Completed"
	// HyperJobFailed means the job has failed its execution.
	HyperJobFailed HyperJobConditionType = "Failed"
	// HyperJobSuspended means the job is suspended
	HyperJobSuspended HyperJobConditionType = "Suspended"
	// HyperJobStartupPolicyCompleted means the StartupPolicy was complete
	HyperJobStartupPolicyCompleted HyperJobConditionType = "StartupPolicyCompleted"
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
	// MaxRestarts defines the limit on the number of HyperJob restarts.
	// A restart is achieved by recreating all active child jobs.
	MaxRestarts int32 `json:"maxRestarts,omitempty"`
}

type SuccessPolicy struct {
	// Operator determines either All or Any of the selected jobs should succeed to consider the HyperJob successful
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
	// InOrder means to start them as they are listed in the HyperJob. A ReplicatedJob is started only
	// when all the jobs of the previous one are ready.
	// +kubebuilder:validation:Enum=AnyOrder;InOrder
	StartupPolicyOrder StartupPolicyOptions `json:"startupPolicyOrder"`
}
