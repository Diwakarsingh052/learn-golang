package main

import (
	"errors"
	"fmt"
	"os"
)

var fileNotFound error = errors.New("any additional info")

//var err = fmt.Errorf("%v", "i am an error")

func openFile(fileName string) error {
	_, err := os.Open(fileName)

	if err != nil {

		return fmt.Errorf("%v %w", err, fileNotFound) // returns error by wrapping additional info
	}
	return nil
}

func main() {

	err := openFile("any.txt")

	if errors.Is(err, fileNotFound) { // does err contains fileNotFound or not
		fmt.Println("match")
	} else {
		fmt.Println("not")
	}


}
