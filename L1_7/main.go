package main

import (
	"fmt"
	"sync"
)

func LockMap(wg *sync.WaitGroup, i int, mx *sync.Mutex, myMap map[int]int) { //map это ссылочный тип, поэтому могу просто в функцию передавать и изменения будут и снаружи
	defer wg.Done()
	mx.Lock()
	myMap[1] += i
	mx.Unlock()
}

func main() {
	testMap := make(map[int]int)
	var wg sync.WaitGroup
	var mx sync.Mutex
	wg.Add(15)
	sum := 0
	for i := range 15 {
		go LockMap(&wg, i, &mx, testMap)
		sum += i
	}
	wg.Wait()
	fmt.Println(testMap[1], sum)

}
