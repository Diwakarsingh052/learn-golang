package main

import (
	"encoding/json"
	"log"
	"os"
)

type Perms map[string]bool // ["admin"] true / false
type student struct {
	FirstName string `json:"first_name"` //when you are sending data to a network we don't use caps names
	Password  string `json:"-"`          // omit this field
	Perms     `json:"perms,omitempty"`   // embedding only happens with a custom type
}

func main() {
	stu := []student{
		{
			FirstName: "Ajay",
			Password:  "abc1234",
			Perms:     Perms{"admin": true},
		},
		{
			FirstName: "Rahul",
			Password:  "xyz",
		},
		{
			FirstName: "John",
			Password:  "abc1234",
			Perms:     Perms{"write": false},
		},
	}

	//jsonData, err := json.Marshal(stu)
	jsonData, err := json.MarshalIndent(stu, "", "\t") // don't need to use in project
	if err != nil {
		log.Fatal(err)
	}
	//ioutil // never use this // deprecated in go 1.16
	err = os.WriteFile("user.json", jsonData, 666) // octal permission // refer octal.txt
	if err != nil {
		log.Fatal(err)
	}
}
