package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	numRoutine := 2
	wg.Add(numRoutine)

	for i := 0; i < numRoutine; i++ {
		go print()
	}
	wg.Wait()

}

func print() {
	fmt.Println("This is a goroutine")
	wg.Done()
}
