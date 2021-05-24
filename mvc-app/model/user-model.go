package model

import (
	"fmt"
	"github.com/username/repoName/utils"
	"net/http"
)

var users = map[uint64]*User{
	123: &User{
		FName: "Raj",
		LName: "Ahuja",
		Email: "raj@email.com",
	},
	124: &User{
		FName: "Dev",
		LName: "Kumar",
		Email: "dev@email.com",
	},
}

//GetUser accepts an userId and return the user if exists or an error if not
func GetUser(userId uint64) (*User, *utils.ApplicationError) {
	u := users[userId]
	// user present -> data , not nil
	// userId not found -> nil

	if u != nil { // data present
		return u, nil // stop the func and return the results
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User not found with user id %v", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
