// +build !ignore_autogenerated

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&ConfigMap{}, func(obj interface{}) { SetObjectDefaults_ConfigMap(obj.(*ConfigMap)) })
	scheme.AddTypeDefaultingFunc(&ConfigMapList{}, func(obj interface{}) { SetObjectDefaults_ConfigMapList(obj.(*ConfigMapList)) })
	scheme.AddTypeDefaultingFunc(&Namespace{}, func(obj interface{}) { SetObjectDefaults_Namespace(obj.(*Namespace)) })
	scheme.AddTypeDefaultingFunc(&NamespaceList{}, func(obj interface{}) { SetObjectDefaults_NamespaceList(obj.(*NamespaceList)) })
	scheme.AddTypeDefaultingFunc(&Repository{}, func(obj interface{}) { SetObjectDefaults_Repository(obj.(*Repository)) })
	scheme.AddTypeDefaultingFunc(&RepositoryList{}, func(obj interface{}) { SetObjectDefaults_RepositoryList(obj.(*RepositoryList)) })
	return nil
}

func SetObjectDefaults_ConfigMap(in *ConfigMap) {
	SetDefaults_ConfigMap(in)
}

func SetObjectDefaults_ConfigMapList(in *ConfigMapList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ConfigMap(a)
	}
}

func SetObjectDefaults_Namespace(in *Namespace) {
	SetDefaults_NamespaceSpec(&in.Spec)
}

func SetObjectDefaults_NamespaceList(in *NamespaceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Namespace(a)
	}
}

func SetObjectDefaults_Repository(in *Repository) {
	SetDefaults_RepositorySpec(&in.Spec)
	SetDefaults_RepositoryStatus(&in.Status)
}

func SetObjectDefaults_RepositoryList(in *RepositoryList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Repository(a)
	}
}
