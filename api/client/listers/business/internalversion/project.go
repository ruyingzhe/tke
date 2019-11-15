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
	business "tkestack.io/tke/api/business"
)

// ProjectLister helps list Projects.
type ProjectLister interface {
	// List lists all Projects in the indexer.
	List(selector labels.Selector) (ret []*business.Project, err error)
	// Get retrieves the Project from the index for a given name.
	Get(name string) (*business.Project, error)
	ProjectListerExpansion
}

// projectLister implements the ProjectLister interface.
type projectLister struct {
	indexer cache.Indexer
}

// NewProjectLister returns a new ProjectLister.
func NewProjectLister(indexer cache.Indexer) ProjectLister {
	return &projectLister{indexer: indexer}
}

// List lists all Projects in the indexer.
func (s *projectLister) List(selector labels.Selector) (ret []*business.Project, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*business.Project))
	})
	return ret, err
}

// Get retrieves the Project from the index for a given name.
func (s *projectLister) Get(name string) (*business.Project, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(business.Resource("project"), name)
	}
	return obj.(*business.Project), nil
}
