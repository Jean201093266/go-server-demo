package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"grpc_demo/config"
	"log"
)

var (
	NamingClient naming_client.INamingClient
	ConfigClient config_client.IConfigClient
)

func init() {
	log.Println("init nacos")
	var err error
	nc := config.MyConfig.Nacos

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nc.Ip, nc.Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId: nc.NamespaceId,
		TimeoutMs:   nc.TimeoutMs,
		Username:    nc.Username,
		Password:    nc.Password,
	}
	NamingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		})
	if err != nil {
		panic(err)
	}
	ConfigClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		})
	if err != nil {
		panic(err)
	}
}
