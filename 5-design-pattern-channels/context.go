package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	duration := 1 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)
	var wg = sync.WaitGroup{}
	wg.Add(1)

	go func() {

		time.Sleep(2 * time.Second)
		ch <- "paper"
		wg.Done()

	}()

	select {
	case p := <-ch:
		fmt.Println("work complete", p)
	case <-ctx.Done():
		fmt.Println("Move on")

	}
	wg.Wait()

}
