package config

import (
	"github.com/spf13/viper"
	"log"
)

type serverConfig struct {
	Network                    string
	Address                    string
	Region                     int
	ShouldReportToPrometheus   bool
	ShouldReportRuntimeMetrics bool
}

type mysqlConfig struct {
	DSN             string
	ConnMaxLifetime int
	MaxOpenConns    int
	MaxIdleConns    int
}

type redisConfig struct {
	Addr           string
	Password       string
	DB             int
	IsClusterMode  bool
	ReadOnly       bool
	ClusterAddr    []string
	RouteByLatency bool
	RouteRandomly  bool
}

type nacosConfig struct {
	Ip                  string
	Port                uint64
	NamespaceId         string
	TimeoutMs           uint64
	NotLoadCacheAtStart bool
	Username            string
	Password            string
	Group               string
}

type myServerConfig struct {
	Server serverConfig
	MySQL  mysqlConfig
	Redis  redisConfig
	Nacos  nacosConfig
}

var MyConfig = new(myServerConfig)

func init() {
	log.Println("init config")
	vip := viper.New()
	vip.SetConfigFile("config/config.yml")
	if err := vip.ReadInConfig(); err != nil {
		log.Println(err)
	}
	if err := vip.Unmarshal(MyConfig); err != nil {
		log.Fatalf("init server config error %v", err)
	}
}
