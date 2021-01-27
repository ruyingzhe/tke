/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2021 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 *
 */

package util

import (
	"context"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	"tkestack.io/tke/pkg/platform/util"
)

const (
	// ProjectDatabaseName defines database name for project metrics
	ProjectDatabaseName = "projects"
)

// ClusterNameToClient mapping cluster to kubernetes client
// clusterName => kubernetes.Interface
var ClusterNameToClient sync.Map

// ClusterNameToMonitor mapping cluster to monitoring client
// clusterName => monitoringclient.Clientset
var ClusterNameToMonitor sync.Map

// GetClusterClient get kubernetes client via cluster name
func GetClusterClient(ctx context.Context, clusterName string, platformClient platformversionedclient.PlatformV1Interface) (kubernetes.Interface, error) {
	// First check from cache
	if item, ok := ClusterNameToClient.Load(clusterName); ok {
		// Check if is available
		kubeClient := item.(kubernetes.Interface)
		_, err := kubeClient.CoreV1().Services(metav1.NamespaceSystem).List(ctx, metav1.ListOptions{})
		if err == nil {
			return kubeClient, nil
		}
		ClusterNameToClient.Delete(clusterName)
	}

	kubeClient, err := util.BuildExternalClientSetWithName(ctx, platformClient, clusterName)
	if err != nil {
		return nil, err
	}

	ClusterNameToClient.Store(clusterName, kubeClient)

	return kubeClient, nil
}
