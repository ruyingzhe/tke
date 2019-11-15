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

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	rest "k8s.io/client-go/rest"
	"tkestack.io/tke/api/client/clientset/internalversion/scheme"
)

type BusinessInterface interface {
	RESTClient() rest.Interface
	ConfigMapsGetter
	ImageNamespacesGetter
	NamespacesGetter
	PlatformsGetter
	PortalsGetter
	ProjectsGetter
}

// BusinessClient is used to interact with features provided by the business.tkestack.io group.
type BusinessClient struct {
	restClient rest.Interface
}

func (c *BusinessClient) ConfigMaps() ConfigMapInterface {
	return newConfigMaps(c)
}

func (c *BusinessClient) ImageNamespaces(namespace string) ImageNamespaceInterface {
	return newImageNamespaces(c, namespace)
}

func (c *BusinessClient) Namespaces(namespace string) NamespaceInterface {
	return newNamespaces(c, namespace)
}

func (c *BusinessClient) Platforms() PlatformInterface {
	return newPlatforms(c)
}

func (c *BusinessClient) Portals() PortalInterface {
	return newPortals(c)
}

func (c *BusinessClient) Projects() ProjectInterface {
	return newProjects(c)
}

// NewForConfig creates a new BusinessClient for the given config.
func NewForConfig(c *rest.Config) (*BusinessClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BusinessClient{client}, nil
}

// NewForConfigOrDie creates a new BusinessClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BusinessClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BusinessClient for the given RESTClient.
func New(c rest.Interface) *BusinessClient {
	return &BusinessClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("business.tkestack.io")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("business.tkestack.io")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BusinessClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
