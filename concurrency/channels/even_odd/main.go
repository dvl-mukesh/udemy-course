package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)

	quitch := make(chan struct{})

	go send(eve, odd, quitch)
	receive(eve, odd, quitch)
}

func send(eve, odd chan<- int, quitch chan struct{}) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quitch)
}

func receive(eve, odd <-chan int, quitch chan struct{}) {
	for {
		select {
		case v := <-eve:
			fmt.Printf("Even: %d\n", v)
		case v := <-odd:
			fmt.Printf("odd: %d\n", v)
		case <-quitch:
			return
		}
	}
}
