/*
Copyright 2021 Antrea Authors.

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
	v1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceImportFilterLister helps list ResourceImportFilters.
// All objects returned here must be treated as read-only.
type ResourceImportFilterLister interface {
	// List lists all ResourceImportFilters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceImportFilter, err error)
	// ResourceImportFilters returns an object that can list and get ResourceImportFilters.
	ResourceImportFilters(namespace string) ResourceImportFilterNamespaceLister
	ResourceImportFilterListerExpansion
}

// resourceImportFilterLister implements the ResourceImportFilterLister interface.
type resourceImportFilterLister struct {
	indexer cache.Indexer
}

// NewResourceImportFilterLister returns a new ResourceImportFilterLister.
func NewResourceImportFilterLister(indexer cache.Indexer) ResourceImportFilterLister {
	return &resourceImportFilterLister{indexer: indexer}
}

// List lists all ResourceImportFilters in the indexer.
func (s *resourceImportFilterLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceImportFilter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceImportFilter))
	})
	return ret, err
}

// ResourceImportFilters returns an object that can list and get ResourceImportFilters.
func (s *resourceImportFilterLister) ResourceImportFilters(namespace string) ResourceImportFilterNamespaceLister {
	return resourceImportFilterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceImportFilterNamespaceLister helps list and get ResourceImportFilters.
// All objects returned here must be treated as read-only.
type ResourceImportFilterNamespaceLister interface {
	// List lists all ResourceImportFilters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceImportFilter, err error)
	// Get retrieves the ResourceImportFilter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceImportFilter, error)
	ResourceImportFilterNamespaceListerExpansion
}

// resourceImportFilterNamespaceLister implements the ResourceImportFilterNamespaceLister
// interface.
type resourceImportFilterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceImportFilters in the indexer for a given namespace.
func (s resourceImportFilterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceImportFilter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceImportFilter))
	})
	return ret, err
}

// Get retrieves the ResourceImportFilter from the indexer for a given namespace and name.
func (s resourceImportFilterNamespaceLister) Get(name string) (*v1alpha1.ResourceImportFilter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourceimportfilter"), name)
	}
	return obj.(*v1alpha1.ResourceImportFilter), nil
}