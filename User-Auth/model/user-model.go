package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}



func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})

	if err != nil {
		return nil , err
	}



}
