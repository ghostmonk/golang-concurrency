package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)

	go func(wg *sync.WaitGroup, ch <-chan int) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}

		wg.Done()
	}(wg, ch)

	go func(wg *sync.WaitGroup, ch chan<- int) {
		ch <- 32
		close(ch)
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
