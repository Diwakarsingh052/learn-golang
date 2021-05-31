package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// I pass a value to my server and it returns a double of that
func main() {
	http.HandleFunc("/double", doubleHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("v") // trying to fetch value from the user
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, v*2)
}
