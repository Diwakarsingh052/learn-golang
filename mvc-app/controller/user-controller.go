package controller

import (
	"encoding/json"
	"fmt"
	"github.com/username/repoName/model"
	"github.com/username/repoName/utils"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//panic("panic")
	fmt.Fprintln(w, "This is our Home Page")

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	userIdString := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		appErr := utils.ApplicationError{
			Message:    "User Id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(appErr)
		w.WriteHeader(appErr.StatusCode)
		w.Write(jsonValue)
		return

	}
	user, appErr := model.GetUser(userId)
	if appErr != nil {
		jsonValue, _ := json.Marshal(appErr)
		w.WriteHeader(appErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonValue)
	return

}
