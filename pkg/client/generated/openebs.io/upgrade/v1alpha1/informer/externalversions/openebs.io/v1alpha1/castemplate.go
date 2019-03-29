/*
Copyright 2019 The OpenEBS Authors

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

	openebsiov1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	internalclientset "github.com/openebs/maya/pkg/client/generated/openebs.io/upgrade/v1alpha1/clientset/internalclientset"
	internalinterfaces "github.com/openebs/maya/pkg/client/generated/openebs.io/upgrade/v1alpha1/informer/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/maya/pkg/client/generated/openebs.io/upgrade/v1alpha1/lister/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CASTemplateInformer provides access to a shared informer and lister for
// CASTemplates.
type CASTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CASTemplateLister
}

type cASTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCASTemplateInformer constructs a new informer for CASTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCASTemplateInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCASTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCASTemplateInformer constructs a new informer for CASTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCASTemplateInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CASTemplates().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CASTemplates().Watch(options)
			},
		},
		&openebsiov1alpha1.CASTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *cASTemplateInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCASTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cASTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openebsiov1alpha1.CASTemplate{}, f.defaultInformer)
}

func (f *cASTemplateInformer) Lister() v1alpha1.CASTemplateLister {
	return v1alpha1.NewCASTemplateLister(f.Informer().GetIndexer())
}