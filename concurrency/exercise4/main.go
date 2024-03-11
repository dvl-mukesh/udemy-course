package main

import (
	"fmt"
	"sync"
)

var counter int

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go increment()
	}

	wg.Wait()

	fmt.Println(counter)
}

func increment() {
	defer mu.Unlock()

	mu.Lock()
	v := counter
	v++
	counter = v
	fmt.Println(counter)
	wg.Done()
}
