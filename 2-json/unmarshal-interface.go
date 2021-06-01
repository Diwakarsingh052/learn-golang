package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var jsonData =  `{
  "users": [
    {
      "name": "Elliot",
      "type": "Reader",
      "age": 23,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    },
    {
      "name": "Fraser",
      "type": "Author",
      "age": 17,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    }
  ]
}`

func main() {

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

}
