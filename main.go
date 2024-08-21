/*
 *
 * Copyright 2015 gRPC authors.
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
 *
 */

// Package grpc_demo implements a server for Greeter service.
package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"grpc_demo/config"
	"grpc_demo/kitex_gen/example/exampleservice"
	"grpc_demo/nacos"
	"log"
)

func main() {
	serverConfig := config.MyConfig.Server
	netAddr := utils.NewNetAddr(serverConfig.Network, serverConfig.Address)

	exampleService := exampleservice.NewServer(
		&ExampleServiceImpl{},
		server.WithServiceAddr(netAddr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "example"}),
		server.WithRegistry(
			registry.NewNacosRegistry(nacos.NamingClient, registry.WithGroup(config.MyConfig.Nacos.Group))))
	if err := exampleService.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
