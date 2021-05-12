package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	Name        string `json:"first_name"`
	Permissions map[string]bool `json:"perms"`
}

func main() {

	data, err := os.ReadFile("user.json")
	if err != nil {
		log.Fatal(err)
	}
	var persons []person

	err = json.Unmarshal(data, &persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(persons)

}
