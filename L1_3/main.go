package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		_ = fmt.Errorf("Error of scanning N %v", err)
	}
	channel := make(chan int, N)

	for i := 0; i < N; i++ { //запуск воркеров
		go func() {
			for {
				apperand := <-channel
				fmt.Println("read from channel ", apperand)
			}
		}()
	}

	for { //бесконечная запись в канал в главной горутине
		channel <- rand.Int()
		time.Sleep(1 * time.Second) // для удобства
		fmt.Println("pause for comfort")

	}

}
