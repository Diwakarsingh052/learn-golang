package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	user := map[string]string{"first_name": "Ajay1", "last_name": "Kumar"}
	jsonValue, _ := json.Marshal(user)
	request, err := http.NewRequest(http.MethodPost, "http://httpbin.org/post", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(data))
}
