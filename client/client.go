package client

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"grpc_demo/config"
	"grpc_demo/kitex_gen/user_service/userservices"
	"grpc_demo/nacos"
	"log"
	"sync"
	"time"
)

var (
	userService userservices.Client
	once        sync.Once
)

func GetUserServiceClient() userservices.Client {
	once.Do(func() {
		var err error
		userService, err = userservices.NewClient(
			"user-service",
			client.WithResolver(
				resolver.NewNacosResolver(
					nacos.NamingClient,
					resolver.WithGroup(config.MyConfig.Nacos.Group))),
			client.WithRPCTimeout(time.Second*3),
			client.WithTransportProtocol(transport.GRPC),
		)
		if err != nil {
			log.Printf("get user service client error %v", err)
		}
	})
	return userService
}
