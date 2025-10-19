package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {
	chFirst := make(chan int, 100)
	chSecond := make(chan int, 100)
	var arr [100]int
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int() % 100
	}

	go func() {
		defer wg.Done()
		defer close(chFirst)
		for _, val := range arr {
			chFirst <- val
		}

	}()

	go func() {
		defer wg.Done()
		defer close(chSecond)
		for val := range chFirst {
			chSecond <- val * 2
		}
	}()

	go func() {
		defer wg.Done()
		for val := range chSecond {
			fmt.Println(val)
		}
	}()

	wg.Wait()
	fmt.Println("end")

}
