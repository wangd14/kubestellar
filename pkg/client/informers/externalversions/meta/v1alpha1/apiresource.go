//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/meta/v1alpha1"
	scopedclientset "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned"
	clientset "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned/cluster"
	"github.com/kcp-dev/edge-mc/pkg/client/informers/externalversions/internalinterfaces"
	metav1alpha1listers "github.com/kcp-dev/edge-mc/pkg/client/listers/meta/v1alpha1"
)

// APIResourceClusterInformer provides access to a shared informer and lister for
// APIResources.
type APIResourceClusterInformer interface {
	Cluster(logicalcluster.Name) APIResourceInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() metav1alpha1listers.APIResourceClusterLister
}

type aPIResourceClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAPIResourceClusterInformer constructs a new informer for APIResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIResourceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIResourceClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIResourceClusterInformer constructs a new informer for APIResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIResourceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetaV1alpha1().APIResources().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetaV1alpha1().APIResources().Watch(context.TODO(), options)
			},
		},
		&metav1alpha1.APIResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIResourceClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredAPIResourceClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *aPIResourceClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&metav1alpha1.APIResource{}, f.defaultInformer)
}

func (f *aPIResourceClusterInformer) Lister() metav1alpha1listers.APIResourceClusterLister {
	return metav1alpha1listers.NewAPIResourceClusterLister(f.Informer().GetIndexer())
}

// APIResourceInformer provides access to a shared informer and lister for
// APIResources.
type APIResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() metav1alpha1listers.APIResourceLister
}

func (f *aPIResourceClusterInformer) Cluster(clusterName logicalcluster.Name) APIResourceInformer {
	return &aPIResourceInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type aPIResourceInformer struct {
	informer cache.SharedIndexInformer
	lister   metav1alpha1listers.APIResourceLister
}

func (f *aPIResourceInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *aPIResourceInformer) Lister() metav1alpha1listers.APIResourceLister {
	return f.lister
}

type aPIResourceScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func (f *aPIResourceScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&metav1alpha1.APIResource{}, f.defaultInformer)
}

func (f *aPIResourceScopedInformer) Lister() metav1alpha1listers.APIResourceLister {
	return metav1alpha1listers.NewAPIResourceLister(f.Informer().GetIndexer())
}

// NewAPIResourceInformer constructs a new informer for APIResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIResourceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIResourceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIResourceInformer constructs a new informer for APIResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIResourceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetaV1alpha1().APIResources().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetaV1alpha1().APIResources().Watch(context.TODO(), options)
			},
		},
		&metav1alpha1.APIResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIResourceScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIResourceInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}
