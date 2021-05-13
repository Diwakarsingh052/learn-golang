package main

import (
	"fmt"
	"time"
)

func main() { // main is a goroutine with id 1

	go disp() // go creates a goroutine // main by default doesn't wait for a goroutine
	time.Sleep(5 * time.Second) // sleep is unproductive and cpu doesn't wait for it
	// sleep would force cpu to switch to next goroutine ready to run
	fmt.Println("main")
}

func disp() {

	fmt.Println("Hello from disp")
}
