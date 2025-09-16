package main

import (
	"fmt"
	"sync"
)

func Quater(operand int, wg *sync.WaitGroup) {
	fmt.Println(operand * operand)
	wg.Done()
}
func main() {
	data := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(5)

	for i := range 5 {
		go Quater(data[i], &wg)

	}
	wg.Wait()

}
