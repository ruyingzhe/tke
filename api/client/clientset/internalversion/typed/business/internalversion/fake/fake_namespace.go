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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	business "tkestack.io/tke/api/business"
)

// FakeNamespaces implements NamespaceInterface
type FakeNamespaces struct {
	Fake *FakeBusiness
	ns   string
}

var namespacesResource = schema.GroupVersionResource{Group: "business.tkestack.io", Version: "", Resource: "namespaces"}

var namespacesKind = schema.GroupVersionKind{Group: "business.tkestack.io", Version: "", Kind: "Namespace"}

// Get takes name of the namespace, and returns the corresponding namespace object, and an error if there is any.
func (c *FakeNamespaces) Get(name string, options v1.GetOptions) (result *business.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(namespacesResource, c.ns, name), &business.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*business.Namespace), err
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors.
func (c *FakeNamespaces) List(opts v1.ListOptions) (result *business.NamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(namespacesResource, namespacesKind, c.ns, opts), &business.NamespaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &business.NamespaceList{ListMeta: obj.(*business.NamespaceList).ListMeta}
	for _, item := range obj.(*business.NamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namespaces.
func (c *FakeNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(namespacesResource, c.ns, opts))

}

// Create takes the representation of a namespace and creates it.  Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Create(namespace *business.Namespace) (result *business.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(namespacesResource, c.ns, namespace), &business.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*business.Namespace), err
}

// Update takes the representation of a namespace and updates it. Returns the server's representation of the namespace, and an error, if there is any.
func (c *FakeNamespaces) Update(namespace *business.Namespace) (result *business.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(namespacesResource, c.ns, namespace), &business.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*business.Namespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNamespaces) UpdateStatus(namespace *business.Namespace) (*business.Namespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(namespacesResource, "status", c.ns, namespace), &business.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*business.Namespace), err
}

// Delete takes name of the namespace and deletes it. Returns an error if one occurs.
func (c *FakeNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(namespacesResource, c.ns, name), &business.Namespace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(namespacesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &business.NamespaceList{})
	return err
}

// Patch applies the patch and returns the patched namespace.
func (c *FakeNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *business.Namespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(namespacesResource, c.ns, name, pt, data, subresources...), &business.Namespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*business.Namespace), err
}
