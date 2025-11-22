package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) increment() {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) value() {
	fmt.Println("counter >> ", c.count)
}

func main() {
	var wg sync.WaitGroup
	testCounter := Counter{count: 0}

	for range 200 {
		wg.Add(1)
		go func(Count *Counter) {
			defer wg.Done()
			time.Sleep(time.Second)
			Count.increment()
			Count.increment()
		}(&testCounter)
	}

	wg.Wait()

	testCounter.value()
}
