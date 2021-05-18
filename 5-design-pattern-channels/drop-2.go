package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://golang.org",
	"https://github.com",
	"https://gmail.com",
}

func main() {
	ch := make(chan string, 1)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	go func() {
		client := http.Client{}

		for r := range ch {
			resp, _ := client.Get(r)
			fmt.Println("Url", r, "response", resp.Status)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {

		for _, u := range urls {

			select {
			case ch <- u: // []  size 1 -> 2 urls at same time // not possible
				fmt.Println("sent url", u, "to process")
			default:
				fmt.Println("drop and move on", u)

			}

		}
		close(ch) //channel should be closed by that go routine only who sends data
		wg.Done()
	}()


	wg.Wait()
}
