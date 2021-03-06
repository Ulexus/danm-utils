/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	time "time"

	netpolv1 "github.com/nokia/danm-utils/crd/api/netpol/v1"
	versioned "github.com/nokia/danm-utils/crd/client/clientset/versioned"
	internalinterfaces "github.com/nokia/danm-utils/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/nokia/danm-utils/crd/client/listers/netpol/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DanmNetworkPolicyInformer provides access to a shared informer and lister for
// DanmNetworkPolicies.
type DanmNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DanmNetworkPolicyLister
}

type danmNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDanmNetworkPolicyInformer constructs a new informer for DanmNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDanmNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDanmNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDanmNetworkPolicyInformer constructs a new informer for DanmNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDanmNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetpolV1().DanmNetworkPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetpolV1().DanmNetworkPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&netpolv1.DanmNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *danmNetworkPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDanmNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *danmNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&netpolv1.DanmNetworkPolicy{}, f.defaultInformer)
}

func (f *danmNetworkPolicyInformer) Lister() v1.DanmNetworkPolicyLister {
	return v1.NewDanmNetworkPolicyLister(f.Informer().GetIndexer())
}
