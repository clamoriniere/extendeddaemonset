// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/datadog/extendeddaemonset/pkg/apis/datadoghq/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ExtendedDaemonSetReplicaSetLister helps list ExtendedDaemonSetReplicaSets.
type ExtendedDaemonSetReplicaSetLister interface {
	// List lists all ExtendedDaemonSetReplicaSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ExtendedDaemonSetReplicaSet, err error)
	// ExtendedDaemonSetReplicaSets returns an object that can list and get ExtendedDaemonSetReplicaSets.
	ExtendedDaemonSetReplicaSets(namespace string) ExtendedDaemonSetReplicaSetNamespaceLister
	ExtendedDaemonSetReplicaSetListerExpansion
}

// extendedDaemonSetReplicaSetLister implements the ExtendedDaemonSetReplicaSetLister interface.
type extendedDaemonSetReplicaSetLister struct {
	indexer cache.Indexer
}

// NewExtendedDaemonSetReplicaSetLister returns a new ExtendedDaemonSetReplicaSetLister.
func NewExtendedDaemonSetReplicaSetLister(indexer cache.Indexer) ExtendedDaemonSetReplicaSetLister {
	return &extendedDaemonSetReplicaSetLister{indexer: indexer}
}

// List lists all ExtendedDaemonSetReplicaSets in the indexer.
func (s *extendedDaemonSetReplicaSetLister) List(selector labels.Selector) (ret []*v1alpha1.ExtendedDaemonSetReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExtendedDaemonSetReplicaSet))
	})
	return ret, err
}

// ExtendedDaemonSetReplicaSets returns an object that can list and get ExtendedDaemonSetReplicaSets.
func (s *extendedDaemonSetReplicaSetLister) ExtendedDaemonSetReplicaSets(namespace string) ExtendedDaemonSetReplicaSetNamespaceLister {
	return extendedDaemonSetReplicaSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExtendedDaemonSetReplicaSetNamespaceLister helps list and get ExtendedDaemonSetReplicaSets.
type ExtendedDaemonSetReplicaSetNamespaceLister interface {
	// List lists all ExtendedDaemonSetReplicaSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ExtendedDaemonSetReplicaSet, err error)
	// Get retrieves the ExtendedDaemonSetReplicaSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ExtendedDaemonSetReplicaSet, error)
	ExtendedDaemonSetReplicaSetNamespaceListerExpansion
}

// extendedDaemonSetReplicaSetNamespaceLister implements the ExtendedDaemonSetReplicaSetNamespaceLister
// interface.
type extendedDaemonSetReplicaSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExtendedDaemonSetReplicaSets in the indexer for a given namespace.
func (s extendedDaemonSetReplicaSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ExtendedDaemonSetReplicaSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExtendedDaemonSetReplicaSet))
	})
	return ret, err
}

// Get retrieves the ExtendedDaemonSetReplicaSet from the indexer for a given namespace and name.
func (s extendedDaemonSetReplicaSetNamespaceLister) Get(name string) (*v1alpha1.ExtendedDaemonSetReplicaSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("extendeddaemonsetreplicaset"), name)
	}
	return obj.(*v1alpha1.ExtendedDaemonSetReplicaSet), nil
}
