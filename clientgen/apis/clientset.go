// Code generated by client-gen. DO NOT EDIT.

package apis

import (
	"fmt"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	configmanagementv1 "kpt.dev/configsync/clientgen/apis/typed/configmanagement/v1"
	configsyncv1alpha1 "kpt.dev/configsync/clientgen/apis/typed/configsync/v1alpha1"
	configsyncv1beta1 "kpt.dev/configsync/clientgen/apis/typed/configsync/v1beta1"
	hubv1 "kpt.dev/configsync/clientgen/apis/typed/hub/v1"
	kptv1alpha1 "kpt.dev/configsync/clientgen/apis/typed/kpt.dev/v1alpha1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ConfigmanagementV1() configmanagementv1.ConfigmanagementV1Interface
	ConfigsyncV1alpha1() configsyncv1alpha1.ConfigsyncV1alpha1Interface
	ConfigsyncV1beta1() configsyncv1beta1.ConfigsyncV1beta1Interface
	HubV1() hubv1.HubV1Interface
	KptV1alpha1() kptv1alpha1.KptV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	configmanagementV1 *configmanagementv1.ConfigmanagementV1Client
	configsyncV1alpha1 *configsyncv1alpha1.ConfigsyncV1alpha1Client
	configsyncV1beta1  *configsyncv1beta1.ConfigsyncV1beta1Client
	hubV1              *hubv1.HubV1Client
	kptV1alpha1        *kptv1alpha1.KptV1alpha1Client
}

// ConfigmanagementV1 retrieves the ConfigmanagementV1Client
func (c *Clientset) ConfigmanagementV1() configmanagementv1.ConfigmanagementV1Interface {
	return c.configmanagementV1
}

// ConfigsyncV1alpha1 retrieves the ConfigsyncV1alpha1Client
func (c *Clientset) ConfigsyncV1alpha1() configsyncv1alpha1.ConfigsyncV1alpha1Interface {
	return c.configsyncV1alpha1
}

// ConfigsyncV1beta1 retrieves the ConfigsyncV1beta1Client
func (c *Clientset) ConfigsyncV1beta1() configsyncv1beta1.ConfigsyncV1beta1Interface {
	return c.configsyncV1beta1
}

// HubV1 retrieves the HubV1Client
func (c *Clientset) HubV1() hubv1.HubV1Interface {
	return c.hubV1
}

// KptV1alpha1 retrieves the KptV1alpha1Client
func (c *Clientset) KptV1alpha1() kptv1alpha1.KptV1alpha1Interface {
	return c.kptV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.configmanagementV1, err = configmanagementv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.configsyncV1alpha1, err = configsyncv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.configsyncV1beta1, err = configsyncv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.hubV1, err = hubv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kptV1alpha1, err = kptv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.configmanagementV1 = configmanagementv1.NewForConfigOrDie(c)
	cs.configsyncV1alpha1 = configsyncv1alpha1.NewForConfigOrDie(c)
	cs.configsyncV1beta1 = configsyncv1beta1.NewForConfigOrDie(c)
	cs.hubV1 = hubv1.NewForConfigOrDie(c)
	cs.kptV1alpha1 = kptv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.configmanagementV1 = configmanagementv1.New(c)
	cs.configsyncV1alpha1 = configsyncv1alpha1.New(c)
	cs.configsyncV1beta1 = configsyncv1beta1.New(c)
	cs.hubV1 = hubv1.New(c)
	cs.kptV1alpha1 = kptv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
