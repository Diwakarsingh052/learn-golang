package main

import (
	"fmt"
	"learn-go/str"
	"log"
)

func main() {

	i, err := str.ConvertStrToInt("abc")
	if err != nil {
		log.Fatal(err) // no need to write explicit return
	}
	fmt.Println(i)

}
