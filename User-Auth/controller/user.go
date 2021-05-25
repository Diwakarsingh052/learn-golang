package controller

import (
	"auth/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type jsonData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController struct {
	*model.UserService
}

func NewUserController(us *model.UserService) *UserController {
	return &UserController{UserService: us}
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var u jsonData
	json.Unmarshal(b, &u)

	user := model.User{

		Email:    u.Email,
		Password: u.Password,
	}
	err = uc.CreateUser(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Account Created")
}
