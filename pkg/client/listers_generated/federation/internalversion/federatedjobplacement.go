/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	federation "github.com/marun/federation-v2/pkg/apis/federation"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedJobPlacementLister helps list FederatedJobPlacements.
type FederatedJobPlacementLister interface {
	// List lists all FederatedJobPlacements in the indexer.
	List(selector labels.Selector) (ret []*federation.FederatedJobPlacement, err error)
	// FederatedJobPlacements returns an object that can list and get FederatedJobPlacements.
	FederatedJobPlacements(namespace string) FederatedJobPlacementNamespaceLister
	FederatedJobPlacementListerExpansion
}

// federatedJobPlacementLister implements the FederatedJobPlacementLister interface.
type federatedJobPlacementLister struct {
	indexer cache.Indexer
}

// NewFederatedJobPlacementLister returns a new FederatedJobPlacementLister.
func NewFederatedJobPlacementLister(indexer cache.Indexer) FederatedJobPlacementLister {
	return &federatedJobPlacementLister{indexer: indexer}
}

// List lists all FederatedJobPlacements in the indexer.
func (s *federatedJobPlacementLister) List(selector labels.Selector) (ret []*federation.FederatedJobPlacement, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedJobPlacement))
	})
	return ret, err
}

// FederatedJobPlacements returns an object that can list and get FederatedJobPlacements.
func (s *federatedJobPlacementLister) FederatedJobPlacements(namespace string) FederatedJobPlacementNamespaceLister {
	return federatedJobPlacementNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedJobPlacementNamespaceLister helps list and get FederatedJobPlacements.
type FederatedJobPlacementNamespaceLister interface {
	// List lists all FederatedJobPlacements in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*federation.FederatedJobPlacement, err error)
	// Get retrieves the FederatedJobPlacement from the indexer for a given namespace and name.
	Get(name string) (*federation.FederatedJobPlacement, error)
	FederatedJobPlacementNamespaceListerExpansion
}

// federatedJobPlacementNamespaceLister implements the FederatedJobPlacementNamespaceLister
// interface.
type federatedJobPlacementNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedJobPlacements in the indexer for a given namespace.
func (s federatedJobPlacementNamespaceLister) List(selector labels.Selector) (ret []*federation.FederatedJobPlacement, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedJobPlacement))
	})
	return ret, err
}

// Get retrieves the FederatedJobPlacement from the indexer for a given namespace and name.
func (s federatedJobPlacementNamespaceLister) Get(name string) (*federation.FederatedJobPlacement, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(federation.Resource("federatedjobplacement"), name)
	}
	return obj.(*federation.FederatedJobPlacement), nil
}
