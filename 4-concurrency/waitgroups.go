package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//wg.Add(5) // keeps track of goroutines running
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go hello(i) // 5 goroutines // cannot guarantee order of exec
	}
	wg.Wait()
}

func hello( i int)  {
	defer wg.Done() //decrease the counter
	fmt.Println("Hello",i)

}
