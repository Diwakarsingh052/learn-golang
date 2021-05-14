package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	fmt.Println("Started")
	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	go func() {
		c2 <- "two"
	}()

	go func() {
		c3 <- "three"
	}()

	for i := 1; i <= 3; i++ {
		select { // exec the case where channel is ready
		case a := <-c1: // it's a blocking operation // can block your program forever
			fmt.Println(a)
		case b := <-c2:
			fmt.Println(b)
		case c := <-c3:
			fmt.Println(c)
		}
	}

}

//func(a,b string) { // testing // we don't provide any name
//		fmt.Println(a,b)
//	}("hello","hey") // () call of this func occurs just after it body ends
