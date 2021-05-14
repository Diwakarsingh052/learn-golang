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
	ch := make(chan *HttpResponse)

	responses := []*HttpResponse{}

	client := http.Client{}
	for _, url := range urls {

		go func(url string) { // Number of goroutines 3

			fmt.Println("fetching Url", url)
			resp, err := client.Get(url)
			ch <- &HttpResponse{
				url:      url,
				response: resp,
				err:      err,
			}

		}(url)
	}

	for {

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

		case <-time.After(10 * time.Millisecond):
			fmt.Printf(".")

		}

	}

}

func result(res []*HttpResponse) {
	fmt.Println("Output from result func")
	for _, u := range res {
		if u.response != nil {
			fmt.Println(u.response.Status)
		} else {
			fmt.Println(u.url, "not found")

		}
	}

}

func main() {

	resp := httpGet(urls)
	result(resp)

}
