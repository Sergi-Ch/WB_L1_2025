package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func StopWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop second goroutine...")
			return
		default:
			fmt.Println("some work second goroutine...")
			time.Sleep(time.Second * 3)
		}
	}
}

func StopWithSignal(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop chan with signal...")
			return
		default:
			fmt.Println("fourth goroutine is working...")
			time.Sleep(time.Second * 3)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	ch := make(chan bool)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// 1. Канал уведомления
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("stop first goroutine...")
				return
			default:
				fmt.Println("some work...")
				time.Sleep(time.Second * 3)
			}
		}
	}(&wg)

	// 2. runtime.Goexit
	go func() {
		fmt.Println("simulation of some work third goroutine...")
		time.Sleep(time.Second * 3)
		runtime.Goexit()
	}()

	// 3. Контекст
	go StopWithContext(ctx, &wg)

	// 4. Контекст с таймаутом
	go StopWithContext(ctxWithTimeout, &wg)

	// 5. Сигнал — через общий контекст
	go StopWithSignal(ctx, &wg)

	<-sigCh
	fmt.Println("signal received, stopping...")

	ch <- true
	cancel()

	wg.Wait()
	fmt.Println("all goroutines stopped")
}
