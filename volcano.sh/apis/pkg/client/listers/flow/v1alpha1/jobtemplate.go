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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "volcano.sh/apis/pkg/apis/flow/v1alpha1"
)

// JobTemplateLister helps list JobTemplates.
// All objects returned here must be treated as read-only.
type JobTemplateLister interface {
	// List lists all JobTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobTemplate, err error)
	// JobTemplates returns an object that can list and get JobTemplates.
	JobTemplates(namespace string) JobTemplateNamespaceLister
	JobTemplateListerExpansion
}

// jobTemplateLister implements the JobTemplateLister interface.
type jobTemplateLister struct {
	indexer cache.Indexer
}

// NewJobTemplateLister returns a new JobTemplateLister.
func NewJobTemplateLister(indexer cache.Indexer) JobTemplateLister {
	return &jobTemplateLister{indexer: indexer}
}

// List lists all JobTemplates in the indexer.
func (s *jobTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.JobTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobTemplate))
	})
	return ret, err
}

// JobTemplates returns an object that can list and get JobTemplates.
func (s *jobTemplateLister) JobTemplates(namespace string) JobTemplateNamespaceLister {
	return jobTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JobTemplateNamespaceLister helps list and get JobTemplates.
// All objects returned here must be treated as read-only.
type JobTemplateNamespaceLister interface {
	// List lists all JobTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobTemplate, err error)
	// Get retrieves the JobTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.JobTemplate, error)
	JobTemplateNamespaceListerExpansion
}

// jobTemplateNamespaceLister implements the JobTemplateNamespaceLister
// interface.
type jobTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all JobTemplates in the indexer for a given namespace.
func (s jobTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.JobTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobTemplate))
	})
	return ret, err
}

// Get retrieves the JobTemplate from the indexer for a given namespace and name.
func (s jobTemplateNamespaceLister) Get(name string) (*v1alpha1.JobTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("jobtemplate"), name)
	}
	return obj.(*v1alpha1.JobTemplate), nil
}
