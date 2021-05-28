package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", randomData)
	http.ListenAndServe(":8080", nil)

}

func randomData(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Println("Random Data Handler started")
	defer log.Println("Random Data Handler Ended")

	select {
	case <-time.After(1000 * time.Second):
		fmt.Fprintln(w, "Random Data")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
