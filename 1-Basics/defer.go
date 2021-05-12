package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("test", os.O_WRONLY, 666) // test found
	if err !=nil {
		log.Fatal(err)
	}
	defer f.Close() // guarantee to execute even this function fails

	fmt.Println("Hello")
}
