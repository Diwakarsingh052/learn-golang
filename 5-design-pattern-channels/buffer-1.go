package main

import (
	"fmt"
	"strconv"
)

func main() {

	ch := make(chan string, 1)

	go func() {

		for p := range ch {
			fmt.Println("employee working", p)
		}

	}()

	const work = 10
	for w := 0; w < work; w++ {

		ch <- "paper" + strconv.Itoa(w)
	}

	close(ch)

}
