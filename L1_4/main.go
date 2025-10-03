package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop...")
			return
		default:
			fmt.Println("some work...")
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigs
		fmt.Println("graceful shutdown...")
		cancel()
	}()
	wg.Wait()

}
