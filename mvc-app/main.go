package main

import (
	"fmt"
	"github.com/username/repoName/controller"
	"net/http"
)

func main() {

	http.HandleFunc("/home", controller.Home) // give me func that will handle req coming on /home
	http.HandleFunc("/users", controller.GetUser)
	fmt.Println("Started")
	// starts the server // run forever until you stop
	err := http.ListenAndServe(":8080", nil) // default config // default serve mux

	if err != nil {
		panic(err)
	}

}
