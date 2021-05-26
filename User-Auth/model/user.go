package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"not null;unique"   `
	Password     string `json:"password" gorm:"-"`
	PasswordHash string `json:"password_hash"`
	Token        string `json:"token" gorm:"-"`
	TokenHash    string	`json:"token_hash"`
}
