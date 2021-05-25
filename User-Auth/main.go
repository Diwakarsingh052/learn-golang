package main

import (
	"auth/controller"
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
	us.DestructiveReset()

	userC := controller.NewUserController(us)

	r := mux.NewRouter()
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	http.ListenAndServe(":8080", r)

}
