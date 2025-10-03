package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	ch := make(chan int)
	var N int64
	_, _ = fmt.Scan(&N)
	wg.Add(2)

	go func(ctx context.Context, ch chan<- int, group *sync.WaitGroup, N int64) { //пишет
		defer group.Done()
		timer := time.After(time.Duration(N) * time.Second)
		for {
			select {
			case <-timer:
				cancel()
				return
			case ch <- rand.Int() % 20:
				time.Sleep(time.Second * 2)
			}
		}

	}(ctx, ch, &wg, N)
	go func(ctx context.Context, ch <-chan int, group *sync.WaitGroup) { //читает
		defer group.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop...")
				return
			case value := <-ch:
				fmt.Println(value)
			}
		}
	}(ctx, ch, &wg)
	wg.Wait()
}
