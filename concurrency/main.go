package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

var mu sync.Mutex

func increment() {
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	fmt.Printf("Num CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Num Goroutines: %d\n", runtime.NumGoroutine())
	wg.Add(2)
	go increment()
	go increment()
	fmt.Printf("After Num Goroutines: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
