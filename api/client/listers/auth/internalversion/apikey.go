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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	auth "tkestack.io/tke/api/auth"
)

// APIKeyLister helps list APIKeys.
type APIKeyLister interface {
	// List lists all APIKeys in the indexer.
	List(selector labels.Selector) (ret []*auth.APIKey, err error)
	// Get retrieves the APIKey from the index for a given name.
	Get(name string) (*auth.APIKey, error)
	APIKeyListerExpansion
}

// aPIKeyLister implements the APIKeyLister interface.
type aPIKeyLister struct {
	indexer cache.Indexer
}

// NewAPIKeyLister returns a new APIKeyLister.
func NewAPIKeyLister(indexer cache.Indexer) APIKeyLister {
	return &aPIKeyLister{indexer: indexer}
}

// List lists all APIKeys in the indexer.
func (s *aPIKeyLister) List(selector labels.Selector) (ret []*auth.APIKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*auth.APIKey))
	})
	return ret, err
}

// Get retrieves the APIKey from the index for a given name.
func (s *aPIKeyLister) Get(name string) (*auth.APIKey, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(auth.Resource("apikey"), name)
	}
	return obj.(*auth.APIKey), nil
}
