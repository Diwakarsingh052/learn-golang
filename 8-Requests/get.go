package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("https://loripsum.net/api")

	if err != nil {
		log.Fatalln(err)
	}
	data, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(data))
}
