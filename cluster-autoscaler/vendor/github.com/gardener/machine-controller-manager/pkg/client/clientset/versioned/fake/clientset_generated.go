// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned"
	clusterv1alpha1 "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned/typed/cluster/v1alpha1"
	fakeclusterv1alpha1 "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned/typed/cluster/v1alpha1/fake"
	"k8s.io2/apimachinery/pkg/runtime"
	"k8s.io2/apimachinery/pkg/watch"
	"k8s.io2/client-go/discovery"
	fakediscovery "k8s.io2/client-go/discovery/fake"
	"k8s.io2/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// ClusterV1alpha1 retrieves the ClusterV1alpha1Client
func (c *Clientset) ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface {
	return &fakeclusterv1alpha1.FakeClusterV1alpha1{Fake: &c.Fake}
}

// Cluster retrieves the ClusterV1alpha1Client
func (c *Clientset) Cluster() clusterv1alpha1.ClusterV1alpha1Interface {
	return &fakeclusterv1alpha1.FakeClusterV1alpha1{Fake: &c.Fake}
}