// This file was automatically generated by informer-gen

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	controller_v1alpha1 "k8s.io/bgd/pkg/apis/controller/v1alpha1"
	clientset "k8s.io/bgd/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/bgd/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/bgd/pkg/client/listers_generated/controller/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// BlueGreenDeploymentInformer provides access to a shared informer and lister for
// BlueGreenDeployments.
type BlueGreenDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BlueGreenDeploymentLister
}

type blueGreenDeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBlueGreenDeploymentInformer constructs a new informer for BlueGreenDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBlueGreenDeploymentInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBlueGreenDeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBlueGreenDeploymentInformer constructs a new informer for BlueGreenDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBlueGreenDeploymentInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1alpha1().BlueGreenDeployments(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ControllerV1alpha1().BlueGreenDeployments(namespace).Watch(options)
			},
		},
		&controller_v1alpha1.BlueGreenDeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *blueGreenDeploymentInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBlueGreenDeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *blueGreenDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&controller_v1alpha1.BlueGreenDeployment{}, f.defaultInformer)
}

func (f *blueGreenDeploymentInformer) Lister() v1alpha1.BlueGreenDeploymentLister {
	return v1alpha1.NewBlueGreenDeploymentLister(f.Informer().GetIndexer())
}
