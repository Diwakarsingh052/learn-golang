package model

import "errors"

var (
	// RecordNotFound happens when user is not present in Db
	RecordNotFound = errors.New("please create your account first")
	ErrInvalidPassword = errors.New("please enter valid password")
)
