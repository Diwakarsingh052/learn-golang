package main

import (
	"fmt"
	//"log"
	"logger-api/logger"
	"os"
	"os/signal"
	"time"
)

type device struct {
	problem bool // false
}

// implements io.Writer interface
func (d *device) Write(p []byte) (n int, err error) {

	for d.problem { //d.problem == true
		time.Sleep(time.Second)
	}
	fmt.Println(string(p))
	return len(p), nil

}

func main() {
	var d device
	grs := 10
	//l := log.New(&d, "prefix", 0)
	l := logger.New(&d, grs)
	for i := 0; i < grs; i++ {

		go func(id int) {

			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(20 * time.Millisecond)

			}
		}(i)

	}

	// ctrl + c <- signal disk is full

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt) // ctrl + c

	for {
		<-signChan
		d.problem = !d.problem // !true = false , !false = true

	}

}
