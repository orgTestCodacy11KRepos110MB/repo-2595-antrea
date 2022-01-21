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

// ResourceExportFilterLister helps list ResourceExportFilters.
// All objects returned here must be treated as read-only.
type ResourceExportFilterLister interface {
	// List lists all ResourceExportFilters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceExportFilter, err error)
	// ResourceExportFilters returns an object that can list and get ResourceExportFilters.
	ResourceExportFilters(namespace string) ResourceExportFilterNamespaceLister
	ResourceExportFilterListerExpansion
}

// resourceExportFilterLister implements the ResourceExportFilterLister interface.
type resourceExportFilterLister struct {
	indexer cache.Indexer
}

// NewResourceExportFilterLister returns a new ResourceExportFilterLister.
func NewResourceExportFilterLister(indexer cache.Indexer) ResourceExportFilterLister {
	return &resourceExportFilterLister{indexer: indexer}
}

// List lists all ResourceExportFilters in the indexer.
func (s *resourceExportFilterLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceExportFilter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceExportFilter))
	})
	return ret, err
}

// ResourceExportFilters returns an object that can list and get ResourceExportFilters.
func (s *resourceExportFilterLister) ResourceExportFilters(namespace string) ResourceExportFilterNamespaceLister {
	return resourceExportFilterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceExportFilterNamespaceLister helps list and get ResourceExportFilters.
// All objects returned here must be treated as read-only.
type ResourceExportFilterNamespaceLister interface {
	// List lists all ResourceExportFilters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceExportFilter, err error)
	// Get retrieves the ResourceExportFilter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceExportFilter, error)
	ResourceExportFilterNamespaceListerExpansion
}

// resourceExportFilterNamespaceLister implements the ResourceExportFilterNamespaceLister
// interface.
type resourceExportFilterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceExportFilters in the indexer for a given namespace.
func (s resourceExportFilterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceExportFilter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceExportFilter))
	})
	return ret, err
}

// Get retrieves the ResourceExportFilter from the indexer for a given namespace and name.
func (s resourceExportFilterNamespaceLister) Get(name string) (*v1alpha1.ResourceExportFilter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourceexportfilter"), name)
	}
	return obj.(*v1alpha1.ResourceExportFilter), nil
}