package main

import (
	"auth/controller"
	"auth/middleware"
	"auth/model"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// delve
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dsn := "host=localhost user=" + os.Getenv("Name") + " password=" + os.Getenv("Password") + " dbname=" + os.Getenv("database") + " port=" + os.Getenv("Port") + " sslmode=disable"
	us, err := model.NewUserService(dsn)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	//us.DestructiveReset()


	userC := controller.NewUserController(us)
	requireUser := middleware.RequireUser{UserServiceInterface: us}

	r := mux.NewRouter()
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	r.HandleFunc("/login", userC.Login).Methods("GET")
	r.HandleFunc("/hello", requireUser.ApplyFn(userC.Hello)).Methods("GET")
	r.HandleFunc("/", userC.Home)
	http.ListenAndServe(":8080", r)

}
