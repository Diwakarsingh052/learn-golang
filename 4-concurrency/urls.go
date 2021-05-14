package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://golang.org",
	"https://github.com",
	"https://123abccv123.com",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func httpGet(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse) // create a channel to store responses

	responses := []*HttpResponse{}

	client := http.Client{} // creates an http client
	for _, url := range urls {

		go func(url string) { // Number of goroutines 3

			fmt.Println("fetching Url", url)
			resp, err := client.Get(url) // do a get request

			ch <- &HttpResponse{ // sending data in channel
				url:      url,
				response: resp,
				err:      err,
			}

		}(url)
	}

	for {
		// waiting for go routines to put data in channel
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			if r.err != nil {
				fmt.Println("with an error", r.err)

			}
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}

		// after 10 millisecond do something
		case <-time.After(10 * time.Millisecond):
			fmt.Printf(".")

		}

	}

}

func result(res []*HttpResponse) {
	fmt.Println("Output from result func")
	for _, u := range res {
		if u.response != nil { // if a domain is not found then response is nil
			fmt.Println(u.url, u.response.Status)
		} else {
			fmt.Println(u.url, "not found")

		}
	}

}

func main() {

	resp := httpGet(urls)
	result(resp)

}
