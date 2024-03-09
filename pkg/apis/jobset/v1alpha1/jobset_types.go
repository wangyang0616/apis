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
	batchv1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

// JobsetSpec defines the desired state of Jobset
type JobsetSpec struct {
	// ReplicatedJobs is a group of volcano jobs managed by jobset.
	// +listType=map
	// +listMapKey=name
	ReplicatedJobs []ReplicatedJob `json:"replicatedJobs,omitempty"`

	// The minimal available jobs to run for this Jobset.
	// The default is the number of all jobs.
	// +optional
	MinAvailable int32 `json:"minAvailable,omitempty" protobuf:"bytes,2,opt,name=minAvailable"`

	// Specifies the default lifecycle of jobs
	// +optional
	// Policies []LifecyclePolicy `json:"policies,omitempty" protobuf:"bytes,5,opt,name=policies"`

	//Specifies the queue that will be used in the scheduler, "default" queue is used this leaves empty.
	// +optional
	Queue string `json:"queue,omitempty" protobuf:"bytes,8,opt,name=queue"`

	// Specifies the maximum number of retries before marking this Jobset failed.
	// Defaults to 3.
	// +optional
	MaxRetry int32 `json:"maxRetry,omitempty" protobuf:"bytes,9,opt,name=maxRetry"`

	// If specified, indicates the job's priority.
	// +optional
	PriorityClassName string `json:"priorityClassName,omitempty" protobuf:"bytes,11,opt,name=priorityClassName"`
}

// JobsetStatus defines the observed state of Jobset
type JobsetStatus struct {
	// Current state of Job.
	// +optional
	State JobsetState `json:"state,omitempty" protobuf:"bytes,1,opt,name=state"`

	// The minimal available pods to run for this Job
	// +optional
	MinAvailable int32 `json:"minAvailable,omitempty" protobuf:"bytes,2,opt,name=minAvailable"`

	// ReplicatedJobsStatus track the number of JobsReady for each replicatedJob.
	// +optional
	// +listType=map
	// +listMapKey=name
	ReplicatedJobsStatus []ReplicatedJobStatus `json:"replicatedJobsStatus,omitempty"`

	//Current version of job
	// +optional
	Version int32 `json:"version,omitempty" protobuf:"bytes,9,opt,name=version"`

	// The number of Job retries.
	// +optional
	RetryCount int32 `json:"retryCount,omitempty" protobuf:"bytes,10,opt,name=retryCount"`

	// Which conditions caused the current Jobset state.
	// +optional
	// +patchMergeKey=status
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=status
	Conditions []JobsetCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"status" protobuf:"bytes,13,rep,name=conditions"`
}

//+genclient
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:resource:path=jobset,shortName=js
//+kubebuilder:subresource:status

// Jobset is the Schema for the jobsets API
type Jobset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JobsetSpec   `json:"spec,omitempty"`
	Status JobsetStatus `json:"status,omitempty"`
}

//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// JobsetList contains a list of Jobset
type JobsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Jobset `json:"items"`
}

type ReplicatedJob struct {
	// Name is the name of the entry and will be used as a suffix for the Job name.
	Name string `json:"name"`
	// Template defines the template of the Job that will be created.
	Template batchv1alpha1.JobSpec `json:"template"`

	// Replicas is the number of jobs that will be created from this ReplicatedJob's template.
	// Jobs names will be in the format: <jobSet.name>-<spec.replicatedJob.name>-<job-index>
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`
}

// JobsetPhase defines the phase of the jobset.
type JobsetPhase string

const (
	// Pending is the phase that jobset is pending in the queue, waiting for scheduling decision
	Pending JobsetPhase = "Pending"
	// Aborting is the phase that jobset is aborted, waiting for releasing jobs
	Aborting JobsetPhase = "Aborting"
	// Aborted is the phase that jobset is aborted by user or error handling
	Aborted JobsetPhase = "Aborted"
	// Running is the phase that minimal available jobs of Job are running
	Running JobsetPhase = "Running"
	// Restarting is the phase that the jobset is restarted, waiting for job releasing and recreating
	Restarting JobsetPhase = "Restarting"
	// Completing is the phase that required jobs are completed, jobset starts to clean up
	Completing JobsetPhase = "Completing"
	// Completed is the phase that all jobs of jobset are completed
	Completed JobsetPhase = "Completed"
	// Terminating is the phase that the jobset is terminated, waiting for releasing jobs
	Terminating JobsetPhase = "Terminating"
	// Terminated is the phase that the jobset is finished unexpected, e.g. events
	Terminated JobsetPhase = "Terminated"
	// Failed is the phase that the jobset is restarted failed reached the maximum number of retries.
	Failed JobsetPhase = "Failed"
)

// JobsetState contains details for the current state of the jobset.
type JobsetState struct {
	// The phase of jobset.
	// +optional
	Phase JobsetPhase `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase"`

	// Unique, one-word, CamelCase reason for the phase's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,2,opt,name=reason"`

	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`

	// Last time the condition transit from one phase to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
}

// ReplicatedJobStatus defines the observed ReplicatedJobs status.
// Statistics of status information of a type of job.
type ReplicatedJobStatus struct {
	Name        string `json:"name"`
	Pending     int32  `json:"pending"`
	Aborting    int32  `json:"aborting"`
	Aborted     int32  `json:"aborted"`
	Running     int32  `json:"running"`
	Restarting  int32  `json:"restarting"`
	Completing  int32  `json:"completing"`
	Completed   int32  `json:"completed"`
	Terminating int32  `json:"terminating"`
	Terminated  int32  `json:"terminated"`
	Failed      int32  `json:"failed"`
}

// JobsetCondition contains details for the current condition of this Jobset.
type JobsetCondition struct {
	// Status is the new phase of jobset after performing the state's action.
	Status JobsetPhase `json:"status" protobuf:"bytes,1,opt,name=status,casttype=JobsetPhase"`

	// Last time the condition transitioned from one phase to another.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,2,opt,name=lastTransitionTime"`
}

func init() {
	SchemeBuilder.Register(&Jobset{}, &JobsetList{})
}
