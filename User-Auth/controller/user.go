package controller

import (
	"auth/model"
	"auth/rand"
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
	err = uc.signIn(w, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Account Created")
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userDetails model.User

	err = json.Unmarshal(b, &userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := uc.Authenticate(userDetails.Email, userDetails.Password)

	if err != nil {
		switch err {
		case model.RecordNotFound:
			http.Error(w, "Invalid email address", http.StatusInternalServerError)
		case model.ErrInvalidPassword:
			http.Error(w, "Invalid Password", http.StatusInternalServerError)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	}

	err = uc.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Logged In", user)

}

func (uc *UserController) signIn(w http.ResponseWriter, user *model.User) error {

	token, err := rand.Token()

	if err != nil {
		return err
	}

	user.Token = token

	err = uc.SaveToken(user) // store in db in hash
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:  "token",
		Value: user.Token,
	}
	http.SetCookie(w, &cookie)
	return nil
}
func (uc *UserController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is our home page")
}

func (uc *UserController) Hello(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("token")
	user, _ := uc.SearchToken(cookie.Value)
	
	fmt.Fprintln(w, "Hello", user.Email)

}
