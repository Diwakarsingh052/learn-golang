package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctx := context.Background()                            // parent context // empty context with no deadlines
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // maximum time allocated for the request
	defer cancel()                                         // cleanup resources
	PrintData(ctx, 2*time.Second, "hello")                 // this func takes n amount of time to process a req

}

func PrintData(ctx context.Context, d time.Duration, msg string) {

	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())

	}

}
