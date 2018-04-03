// This file was automatically generated by informer-gen

package v1alpha2

import (
	internalinterfaces "github.com/openshift/service-idler/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Idlers returns a IdlerInformer.
	Idlers() IdlerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Idlers returns a IdlerInformer.
func (v *version) Idlers() IdlerInformer {
	return &idlerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
