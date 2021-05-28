package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"not null;unique"   `
	Password     string `json:"password" gorm:"-"`
	PasswordHash string `json:"password_hash"`
	Token        string `json:"token" gorm:"-"`
	TokenHash    string `json:"token_hash"`
}

// UserServiceInterface is a collection of methods of userService struct
type UserServiceInterface interface {
	CreateUser(user *User) error
	Authenticate(email, password string) (*User, error)
	ByEmail(email string) (*User, error)
	SaveToken(user *User) error
	SearchToken(token string) (*User, error)
	DestructiveReset() error
	Close() error

}


