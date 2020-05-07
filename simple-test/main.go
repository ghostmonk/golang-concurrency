package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]string{}
var db = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup) {
			if v, ok := queryCache(id); ok {
				fmt.Println("From Cache: " + v)
			}
			wg.Done()
		}(id, wg)

		go func(id int, wg *sync.WaitGroup) {
			if v, ok := queryDatabase(id); ok {
				fmt.Println("From DB: " + v)
				cache[id] = v
			}
			wg.Done()
		}(id, wg)

		wg.Wait()
	}
}

func queryCache(id int) (string, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (string, bool) {
	time.Sleep(100 * time.Millisecond)
	for i, v := range db {
		if id == i {
			return v, true
		}
	}
	return "", false
}
