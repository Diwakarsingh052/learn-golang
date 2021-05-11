package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:] //take up to the last value

	if len(args) < 2 {
		log.Println("Please Provide your name and age")
		return
	}
	fmt.Println(args)
	name := args[0]
	ageString := args[1]

	marksString := args[2]

	age, err := strconv.Atoi(ageString)

	if err != nil {
		log.Fatal(err)
	}

	marks, err := strconv.Atoi(marksString)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(marks)

}