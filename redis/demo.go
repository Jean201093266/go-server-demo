package redis

import (
	"context"
	"log"
)

func Set(key string, value string) error {
	err := GetInstance().client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		log.Printf("Set error %v", err)
	}
	return err
}

func Get(key string) string {
	res, err := GetInstance().client.Get(context.Background(), key).Result()
	if err != nil {
		log.Printf("Set error %v", err)
		return ""
	}
	return res
}
