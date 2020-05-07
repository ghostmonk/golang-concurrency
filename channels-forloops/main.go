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
		for msg := range ch {
			fmt.Println(msg)
		}

		wg.Done()
	}(wg, ch)

	go func(wg *sync.WaitGroup, ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(wg, ch)

	wg.Wait()
}
