package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed value -> basis of it go gen random no.
	// seed value will be fixed by default

	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	fmt.Println(r)

	r = rand.Intn(100)
	fmt.Println(r)
	r = rand.Intn(100)
	fmt.Println(r)
	r = rand.Intn(100)
	fmt.Println(r)

}
