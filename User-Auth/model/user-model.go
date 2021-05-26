package model

import (
	"auth/hash"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserService struct {
	db        *gorm.DB
	hash.HMAC //custom impl
}

const userPWPepper = "secret-random-string"
const secretKey = "secret-random-string"

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	hash := hash.NewHMAC(secretKey)
	u := &UserService{
		db:   db,
		HMAC: hash,
	}
	return u, nil

}

// CreateUser will create the user in DB with email and hashed pass. Handling Signup request from controller
func (us *UserService) CreateUser(user *User) error {

	pwBytes := []byte(user.Password + userPWPepper)
	HashedPass, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = ""

	user.PasswordHash = string(HashedPass)
	err = us.db.Create(user).Error

	if err != nil {
		return err
	}
	return nil

}

// Authenticate will help to log in the user
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email) // checking for email exist or not
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPWPepper))

	switch err {
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrInvalidPassword
	case nil:
		return foundUser, nil
	default:
		return nil, err
	}

}

// ByEmail search users on basis of email
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email=?", email)
	err := first(db, &user)

	if err != nil {

		return nil, err
	}
	return &user, nil
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return RecordNotFound
	case nil:
		return nil
	default:
		return err
	}
}

func (us *UserService) SaveToken(user *User) error {

	user.TokenHash = us.HMAC.Hash(user.Token)
	err := us.db.Model(&User{}).Where("email=?", user.Email).Update("token_hash", user.TokenHash).Error
	return err
}

func (us *UserService) SearchToken(token string) (*User, error) {
	return nil, nil
}

//DestructiveReset drops the table don't use in prod
func (us *UserService) DestructiveReset() error {
	err := us.db.Migrator().DropTable(&User{})
	if err != nil {
		return err
	}
	err = us.db.Migrator().AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Close() error {

	conn, err := us.db.DB()
	if err != nil {
		return err
	}
	conn.Close()
	return nil

}
