package main

import (
	"fmt"
	"sync"
	"time"
)

func or(cs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{}) // Объявление single-канала
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs { // Обход каналов
		go func(c <-chan interface{}) { // Горутина для записи данных в single-канала
			out <- c
			wg.Done()
		}(c)
	}
	go func() { // Горутина для закрытия канала
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}
