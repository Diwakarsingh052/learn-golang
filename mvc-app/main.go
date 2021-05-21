package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/home", home) // give me func that will handle req coming on /home

	fmt.Println("Started")
	// starts the server // run forever until you stop
	err := http.ListenAndServe(":8080", nil) // default config // default serve mux

	if err != nil {
		panic(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	//panic("panic")
	fmt.Fprintln(w, "This is our Home Page")

}
