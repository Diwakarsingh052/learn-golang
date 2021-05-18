package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	const cap = 2

	ch := make(chan string, cap)
	wg.Add(1)
	go func() {
		for p := range ch {
			fmt.Println("emp recv'd", p)
		}
		wg.Done()
	}()

	const work = 5

	for w := 0; w < work; w++ {
		select {

		case ch <- "paper":
			fmt.Println("manager send ack")
		default:
			fmt.Println("manager :drop ")


		}
	}
	close(ch)
	wg.Wait()


	// [t1,t2 ] t3 // t3 drops
	//t2 finish
	//[t1, t4] t4 // t4 accepts
}
