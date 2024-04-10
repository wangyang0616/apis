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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
)

// FakeQueues implements QueueInterface
type FakeQueues struct {
	Fake *FakeSchedulingV1beta1
}

var queuesResource = v1beta1.SchemeGroupVersion.WithResource("queues")

var queuesKind = v1beta1.SchemeGroupVersion.WithKind("Queue")

// Get takes name of the queue, and returns the corresponding queue object, and an error if there is any.
func (c *FakeQueues) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(queuesResource, name), &v1beta1.Queue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Queue), err
}

// List takes label and field selectors, and returns the list of Queues that match those selectors.
func (c *FakeQueues) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.QueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(queuesResource, queuesKind, opts), &v1beta1.QueueList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.QueueList{ListMeta: obj.(*v1beta1.QueueList).ListMeta}
	for _, item := range obj.(*v1beta1.QueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested queues.
func (c *FakeQueues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(queuesResource, opts))
}

// Create takes the representation of a queue and creates it.  Returns the server's representation of the queue, and an error, if there is any.
func (c *FakeQueues) Create(ctx context.Context, queue *v1beta1.Queue, opts v1.CreateOptions) (result *v1beta1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(queuesResource, queue), &v1beta1.Queue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Queue), err
}

// Update takes the representation of a queue and updates it. Returns the server's representation of the queue, and an error, if there is any.
func (c *FakeQueues) Update(ctx context.Context, queue *v1beta1.Queue, opts v1.UpdateOptions) (result *v1beta1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(queuesResource, queue), &v1beta1.Queue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Queue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeQueues) UpdateStatus(ctx context.Context, queue *v1beta1.Queue, opts v1.UpdateOptions) (*v1beta1.Queue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(queuesResource, "status", queue), &v1beta1.Queue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Queue), err
}

// Delete takes name of the queue and deletes it. Returns an error if one occurs.
func (c *FakeQueues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(queuesResource, name, opts), &v1beta1.Queue{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQueues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(queuesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.QueueList{})
	return err
}

// Patch applies the patch and returns the patched queue.
func (c *FakeQueues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Queue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(queuesResource, name, pt, data, subresources...), &v1beta1.Queue{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Queue), err
}
