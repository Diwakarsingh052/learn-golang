package main

import (
	"fmt"
	"sync"
	"time"
)

var cabs = 2
var wg = sync.WaitGroup{}

func main() {

	// mutex
	m := &sync.Mutex{} // if passing to func around make sure to pass a reference

	names := []string{"Ravi", "Raj", "Dev", "Ankit", "Vipin"}

	for _, name := range names {
		wg.Add(1)
		go bookCab(name, m)
	}
	wg.Wait()
}

func bookCab(name string, m *sync.Mutex) {

	defer wg.Done()
	fmt.Println("Welcome to our website", name)
	m.Lock() // Lock the critical section until unlock is not called
	defer m.Unlock()
	//critical section
	if cabs >= 1 {

		fmt.Println("cab is available for", name)
		time.Sleep(1 * time.Second)
		fmt.Println("Booking Confirmed for ", name)
		fmt.Println("Thanks", name)
		cabs--

	} else {
		fmt.Println("cab is not available for", name)

	}
	//m.Unlock()

}
