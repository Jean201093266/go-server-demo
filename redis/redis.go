package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"grpc_demo/config"
	"grpc_demo/nacos"
	"log"
	"sync"
)

type Client struct {
	client         *redis.Client
	options        *redis.Options
	clusterClient  *redis.ClusterClient
	clusterOptions *redis.ClusterOptions
}

type nacosConfig struct {
	SRIDS_HOST     string
	SRIDS_PORT     string
	SRIDS_PASSWORD string
}

var (
	redisClient *Client
	once        sync.Once
)

func GetInstance() *Client {
	once.Do(func() {
		log.Println("init redis")
		redisClient = &Client{}
		content, err := nacos.ConfigClient.GetConfig(
			vo.ConfigParam{
				DataId: "redis_config",
				Group:  config.MyConfig.Nacos.Group,
			})
		if err != nil {
			panic(err)
		}
		rc := initNacosConfig(content)
		redisClient.options = &redis.Options{
			Addr:     fmt.Sprintf("%v:%v", rc.SRIDS_HOST, rc.SRIDS_PORT),
			Password: rc.SRIDS_PASSWORD,
		}
		redisClient.client = redis.NewClient(redisClient.options)
		go func() {
			err := nacos.ConfigClient.ListenConfig(
				vo.ConfigParam{
					DataId: "redis_config",
					Group:  config.MyConfig.Nacos.Group,
					OnChange: func(namespace, group, dataId, data string) {
						log.Println("redis_config changed")
						rc := initNacosConfig(data)
						redisClient.options = &redis.Options{
							Addr:     fmt.Sprintf("%v:%v", rc.SRIDS_HOST, rc.SRIDS_PORT),
							Password: rc.SRIDS_PASSWORD,
						}
						redisClient.client = redis.NewClient(redisClient.options)
					},
				})
			if err != nil {
				log.Printf("listen redis_config error %v", err)
				return
			}
		}()
	})
	return redisClient
}

func initNacosConfig(data string) *nacosConfig {
	rc := &nacosConfig{}
	if err := json.Unmarshal([]byte(data), rc); err != nil {
		log.Printf("init redis_config error %v", err)
	}
	return rc
}
