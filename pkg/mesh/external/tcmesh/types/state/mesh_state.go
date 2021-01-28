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

package state

type MeshState struct {
	Clusters []*ClusterState `json:"clusters"`

	IstiodMetaLB  string `json:"istiod_meta_lb,omitempty"`
	APIServerLB   string `json:"api_server_lb,omitempty"`
	APIServerIP   string `json:"api_server_ip,omitempty"`
	APIServerPort int    `json:"api_server_port,omitempty"`

	// CA
	RootCert string `json:"rootcert"`
	RootKey  string `json:"rootkey"`
}
