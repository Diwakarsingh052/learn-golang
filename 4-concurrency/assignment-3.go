package main

import (
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// a[1,3,67,89] // b := a[1:3] //  slice int

	fileNames := os.Args[1:] // type of FileName -> slice of string

	//fmt.Printf("%T", fileNames)
	//for i:=0 ; i < len(fileNames); i++ { }
	for _, f := range fileNames { // string
		wg.Add(1)
		go createFile(f)
	}

	wg.Wait()

}

func createFile(name string) {

	err := os.WriteFile(name, []byte("Hey"), 666)

	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
}
