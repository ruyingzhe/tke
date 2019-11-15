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
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeLBCFs implements LBCFInterface
type FakeLBCFs struct {
	Fake *FakePlatformV1
}

var lbcfsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "lbcfs"}

var lbcfsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "LBCF"}

// Get takes name of the lBCF, and returns the corresponding lBCF object, and an error if there is any.
func (c *FakeLBCFs) Get(name string, options v1.GetOptions) (result *platformv1.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(lbcfsResource, name), &platformv1.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.LBCF), err
}

// List takes label and field selectors, and returns the list of LBCFs that match those selectors.
func (c *FakeLBCFs) List(opts v1.ListOptions) (result *platformv1.LBCFList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(lbcfsResource, lbcfsKind, opts), &platformv1.LBCFList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.LBCFList{ListMeta: obj.(*platformv1.LBCFList).ListMeta}
	for _, item := range obj.(*platformv1.LBCFList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lBCFs.
func (c *FakeLBCFs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(lbcfsResource, opts))
}

// Create takes the representation of a lBCF and creates it.  Returns the server's representation of the lBCF, and an error, if there is any.
func (c *FakeLBCFs) Create(lBCF *platformv1.LBCF) (result *platformv1.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(lbcfsResource, lBCF), &platformv1.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.LBCF), err
}

// Update takes the representation of a lBCF and updates it. Returns the server's representation of the lBCF, and an error, if there is any.
func (c *FakeLBCFs) Update(lBCF *platformv1.LBCF) (result *platformv1.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(lbcfsResource, lBCF), &platformv1.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.LBCF), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLBCFs) UpdateStatus(lBCF *platformv1.LBCF) (*platformv1.LBCF, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(lbcfsResource, "status", lBCF), &platformv1.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.LBCF), err
}

// Delete takes name of the lBCF and deletes it. Returns an error if one occurs.
func (c *FakeLBCFs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(lbcfsResource, name), &platformv1.LBCF{})
	return err
}

// Patch applies the patch and returns the patched lBCF.
func (c *FakeLBCFs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.LBCF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(lbcfsResource, name, pt, data, subresources...), &platformv1.LBCF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.LBCF), err
}
