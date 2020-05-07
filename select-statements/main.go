package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	dbChan := make(chan Book)
	cacheChan := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, ch chan<- Book) {
			if b, ok := queryCache(id); ok {
				ch <- b	
			}
			wg.Done()
		}(id, wg, cacheChan)

		go func(id int, wg *sync.WaitGroup, ch chan<- Book) {
			if b, ok := queryDatabase(id); ok {
				cache[id] = b
				ch <- b	
			}
			wg.Done()
		}(id, wg, dbChan)

		go func(dbCh, cacheCh <-chan Book) {
			select {
			case b := <- dbCh:
				fmt.Println("From Database")
				fmt.Println(b)
			case b := <- cacheCh:
				fmt.Println("From Cache")
				fmt.Println(b)
				<-dbCh
			}
		}(dbChan, cacheChan)

		wg.Wait()
	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}
