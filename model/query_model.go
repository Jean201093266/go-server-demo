package model

import (
	"log"
)

type UserEntity struct {
	Id    int
	Name  string
	Email string
}

func QueryDemo() ([]UserEntity, error) {
	users := make([]UserEntity, 0)
	err := GetInstance().Table("tv_user").Limit(1).Find(&users).Error
	if err != nil {
		log.Printf("GetMessages failed, err %v\n", err)
		return nil, err
	}
	return users, err
}
