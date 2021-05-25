package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB // nil
}

const userPWPepper = "secret-random-string"

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	u := &UserService{db: db}
	return u, nil

}

// CreateUser will create the user in DB with email and hashed pass
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

func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
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
