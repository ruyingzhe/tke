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

package v1

import (
	rest "k8s.io/client-go/rest"
	v1 "tkestack.io/tke/api/business/v1"
	"tkestack.io/tke/api/client/clientset/versioned/scheme"
)

type BusinessV1Interface interface {
	RESTClient() rest.Interface
	ConfigMapsGetter
	ImageNamespacesGetter
	NamespacesGetter
	PlatformsGetter
	PortalsGetter
	ProjectsGetter
}

// BusinessV1Client is used to interact with features provided by the business.tkestack.io group.
type BusinessV1Client struct {
	restClient rest.Interface
}

func (c *BusinessV1Client) ConfigMaps() ConfigMapInterface {
	return newConfigMaps(c)
}

func (c *BusinessV1Client) ImageNamespaces(namespace string) ImageNamespaceInterface {
	return newImageNamespaces(c, namespace)
}

func (c *BusinessV1Client) Namespaces(namespace string) NamespaceInterface {
	return newNamespaces(c, namespace)
}

func (c *BusinessV1Client) Platforms() PlatformInterface {
	return newPlatforms(c)
}

func (c *BusinessV1Client) Portals() PortalInterface {
	return newPortals(c)
}

func (c *BusinessV1Client) Projects() ProjectInterface {
	return newProjects(c)
}

// NewForConfig creates a new BusinessV1Client for the given config.
func NewForConfig(c *rest.Config) (*BusinessV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BusinessV1Client{client}, nil
}

// NewForConfigOrDie creates a new BusinessV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BusinessV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BusinessV1Client for the given RESTClient.
func New(c rest.Interface) *BusinessV1Client {
	return &BusinessV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BusinessV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
