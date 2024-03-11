package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

var wg sync.WaitGroup

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
	v := counter
	runtime.Gosched()
	v++
	counter = v
	fmt.Println(counter)
	wg.Done()
}
