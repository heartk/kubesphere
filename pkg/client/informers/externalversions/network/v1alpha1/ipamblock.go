/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	networkv1alpha1 "kubesphere.io/kubesphere/pkg/apis/network/v1alpha1"
	versioned "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubesphere.io/kubesphere/pkg/client/listers/network/v1alpha1"
)

// IPAMBlockInformer provides access to a shared informer and lister for
// IPAMBlocks.
type IPAMBlockInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IPAMBlockLister
}

type iPAMBlockInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIPAMBlockInformer constructs a new informer for IPAMBlock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIPAMBlockInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIPAMBlockInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIPAMBlockInformer constructs a new informer for IPAMBlock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIPAMBlockInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().IPAMBlocks().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().IPAMBlocks().Watch(options)
			},
		},
		&networkv1alpha1.IPAMBlock{},
		resyncPeriod,
		indexers,
	)
}

func (f *iPAMBlockInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIPAMBlockInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iPAMBlockInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1alpha1.IPAMBlock{}, f.defaultInformer)
}

func (f *iPAMBlockInformer) Lister() v1alpha1.IPAMBlockLister {
	return v1alpha1.NewIPAMBlockLister(f.Informer().GetIndexer())
}
