package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 1 st method
func main() {

	user := map[string]string{"first_name": "Ajay", "last_name": "Kumar"}
	jsonValue, _ := json.Marshal(user)

	resp, err := http.Post("http://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Fatalln(err)
	}

	data, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(data))
}
