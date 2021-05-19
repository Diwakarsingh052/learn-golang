package main

import (
	"fmt"
	"sync"
	"time"
)

//Read Write Mutex

type cache struct {
	rw   sync.RWMutex ////Read Write Mutex
	wg   sync.WaitGroup
	data map[int]string //shared resource
}

func (c *cache) put(i int) {
	c.rw.Lock()
	fmt.Println("Writing to cache")
	c.data[i] = "Random Data"
	time.Sleep(3 * time.Second)
	c.rw.Unlock()
	c.wg.Done()
}

func (c *cache) read(i int) {
	c.rw.RLock() // No one can write at the time when some go reading
	fmt.Printf("Reading Data %#v\n", c.data[i])
	c.rw.RUnlock()
	c.wg.Done()
}

func main() {
	c := cache{
		rw:   sync.RWMutex{},
		wg:   sync.WaitGroup{},
		data: make(map[int]string),
	}

	for i := 1; i <= 10; i++ {
		c.wg.Add(1)
		go c.put(i)
	}

	for i := 1; i <= 10; i++ {
		c.wg.Add(1)
		go c.read(i)
	}

	c.wg.Wait()

}
