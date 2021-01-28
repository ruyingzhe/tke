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

package app

import (
	genericapiserver "k8s.io/apiserver/pkg/server"
	"tkestack.io/tke/cmd/tke-mesh-api/app/config"
	"tkestack.io/tke/pkg/mesh/apiserver"
)

// CreateServerChain creates the apiservers connected via delegation.
func CreateServerChain(cfg *config.Config) (*genericapiserver.GenericAPIServer, error) {
	apiServerConfig := createAPIServerConfig(cfg)
	apiServer, err := CreateAPIServer(apiServerConfig, genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	apiServer.GenericAPIServer.AddPostStartHookOrDie("start-mesh-api-server-informers",
		func(context genericapiserver.PostStartHookContext) error {
			cfg.VersionedSharedInformerFactory.Start(context.StopCh)
			return nil
		})

	return apiServer.GenericAPIServer, nil
}

// CreateAPIServer creates and wires a workable tke-mesh-api
func CreateAPIServer(
	apiServerConfig *apiserver.Config,
	delegateAPIServer genericapiserver.DelegationTarget,
) (*apiserver.APIServer, error) {
	return apiServerConfig.Complete().New(delegateAPIServer)
}

func createAPIServerConfig(cfg *config.Config) *apiserver.Config {
	return &apiserver.Config{
		GenericConfig: &genericapiserver.RecommendedConfig{
			Config: *cfg.GenericAPIServerConfig,
		},
		ExtraConfig: apiserver.ExtraConfig{
			ServerName:              cfg.ServerName,
			VersionedInformers:      cfg.VersionedSharedInformerFactory,
			StorageFactory:          cfg.StorageFactory,
			APIResourceConfigSource: cfg.StorageFactory.APIResourceConfigSource,
			PlatformClient:          cfg.PlatformClient,
			ClusterClients:          cfg.ClusterClients,
			PrivilegedUsername:      cfg.PrivilegedUsername,
			MeshConfig:              cfg.MeshConfig,
			TCMeshClient:            cfg.TCMeshClient,
		},
	}
}
