package model

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpc_demo/config"
	"grpc_demo/nacos"
	"log"
	"sync"
)

type nacosConfig struct {
	MYSQL_TODESK_HOST_MASTER string
	MYSQL_TODESK_PORT_MASTER string
	MYSQL_TODESK_USER_MASTER string
	MYSQL_TODESK_PASS_MASTER string
	MYSQL_TODESK_NAME_MASTER string
}

var (
	db   *gorm.DB
	once sync.Once
)

func GetInstance() *gorm.DB {
	once.Do(func() {
		log.Println("init mysql")
		content, err := nacos.ConfigClient.GetConfig(
			vo.ConfigParam{
				DataId: "mysql_config",
				Group:  config.MyConfig.Nacos.Group,
			})
		if err != nil {
			panic(err)
		}
		dsn := initNacosConfig(content)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Open db failed, err %v\n", err)
		} else {
			if err != nil {
				log.Printf("Get db instance failed, err %v\n", err)
			}
		}
		go func() {
			err := nacos.ConfigClient.ListenConfig(
				vo.ConfigParam{
					DataId: "mysql_config",
					Group:  config.MyConfig.Nacos.Group,
					OnChange: func(namespace, group, dataId, data string) {
						log.Println("mysql_config changed")
						dsn := initNacosConfig(data)
						var err error
						db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
						if err != nil {
							log.Printf("Open db failed, err %v\n", err)
						} else {
							if err != nil {
								log.Printf("Get db instance failed, err %v\n", err)
							}
						}
					},
				})
			if err != nil {
				log.Printf("listen mysql_config error %v", err)
				return
			}
		}()
	})
	return db
}

func initNacosConfig(data string) string {
	mc := &nacosConfig{}
	if err := json.Unmarshal([]byte(data), mc); err != nil {
		log.Printf("init mysql_config error %v", err)
	}
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		mc.MYSQL_TODESK_USER_MASTER,
		mc.MYSQL_TODESK_PASS_MASTER,
		mc.MYSQL_TODESK_HOST_MASTER,
		mc.MYSQL_TODESK_PORT_MASTER,
		mc.MYSQL_TODESK_NAME_MASTER)
}
