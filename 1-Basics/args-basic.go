package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	//fmt.Println(os.Args)
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) < 3 {
		log.Println("Please provide  your name and age and marks")
		return
	}

	name := args[0]
	ageString := args[1]
	marksString := args[2]

	age, err := strconv.Atoi(ageString) //abc // not valid int

	if err != nil { // err = nil // no error
		// err = msg // error
		log.Fatal(err)
	}

	//marks, err := strconv.Atoi(marksString)
	marks, err := strconv.Atoi(marksString) // 200 // nil
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name, age, marks)

}
