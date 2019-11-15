/*
 * Copyright 2019 THL A29 Limited, a Tencent company.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clientsetinternalversion "tkestack.io/tke/api/client/clientset/internalversion"
	internalinterfaces "tkestack.io/tke/api/client/informers/internalversion/internalinterfaces"
	internalversion "tkestack.io/tke/api/client/listers/platform/internalversion"
	platform "tkestack.io/tke/api/platform"
)

// RegistryInformer provides access to a shared informer and lister for
// Registries.
type RegistryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.RegistryLister
}

type registryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRegistryInformer constructs a new informer for Registry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRegistryInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRegistryInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRegistryInformer constructs a new informer for Registry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRegistryInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Platform().Registries().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Platform().Registries().Watch(options)
			},
		},
		&platform.Registry{},
		resyncPeriod,
		indexers,
	)
}

func (f *registryInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRegistryInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *registryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&platform.Registry{}, f.defaultInformer)
}

func (f *registryInformer) Lister() internalversion.RegistryLister {
	return internalversion.NewRegistryLister(f.Informer().GetIndexer())
}
