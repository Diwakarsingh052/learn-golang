package main

import (
	"fmt"
	"sync"
)

// channels are used to share data between goroutines
// channels are blocking

// [  4 , 20  ,9    ]  // assuming it's a channel


var wg = sync.WaitGroup{}

func main() {

	c := make(chan int) // creating the channel of type int
	wg.Add(4) // add 4 go
	go add(2, 2, c)  // spin up add go
	go sub(6, 2, c) // spin up sub go
	go multiply(3, 3, c) // spin up mult go
	go calc(c) // spin up calc go

	wg.Wait()
}

//3rd chance
func add(a, b int, c chan int) {
	defer wg.Done()
	fmt.Println("exec add")
	sum := a + b
	c <- sum // putting value inside channel

}

// 1st chance
func sub(a, b int, c chan int) {
	defer wg.Done()
	fmt.Println("exec sub")
	sum := a - b
	c <- sum // putting value inside channel

}
// 5th chance
func multiply(a, b int, c chan int) {
	defer wg.Done()
	fmt.Println("exec multiply")
	sum := a * b
	c <- sum // putting value inside channel

}

//2nd chance // 4th chance // 6th chance
func calc(c chan int) {
	defer wg.Done()
	fmt.Println("exec calc")

	x, y, z := <-c, <-c, <-c // receiving values from channel // unless we received 3 values in the channel cpu won't exec the next line of code // blocking this func
	fmt.Println(x + y + z) // sum on screen

}
