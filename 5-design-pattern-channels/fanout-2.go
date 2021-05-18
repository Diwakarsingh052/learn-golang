package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fanoutSem()
}

func fanoutSem() {
	emps := 5

	ch := make(chan string, emps)

	// 0 means default or all processors
	//proc := runtime.GOMAXPROCS(0)

	proc := 2
	sem := make(chan bool, proc)

	for e := 0; e < emps; e++ {
		// ctrl+alt+l
		go func(emp int) {
			sem <- true // block if channel send is not possible
			{
				time.Sleep(time.Duration(2 * time.Second))
				ch <- "paper" + strconv.Itoa(emp)
				fmt.Println("emp set a signal", emp)
			}
			<-sem
		}(e)

	}

	for emps > 0 {

		p := <-ch
		emps--
		fmt.Println("manager recv'd signal", p)
	}

}
