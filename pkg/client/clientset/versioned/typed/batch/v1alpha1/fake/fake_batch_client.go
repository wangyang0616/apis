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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "volcano.sh/apis/pkg/client/clientset/versioned/typed/batch/v1alpha1"
)

type FakeBatchV1alpha1 struct {
	*testing.Fake
}

func (c *FakeBatchV1alpha1) HyperJobs(namespace string) v1alpha1.HyperJobInterface {
	return &FakeHyperJobs{c, namespace}
}

func (c *FakeBatchV1alpha1) Jobs(namespace string) v1alpha1.JobInterface {
	return &FakeJobs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBatchV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
