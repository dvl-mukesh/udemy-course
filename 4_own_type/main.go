package main

import "fmt"

type custom int

var x custom

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
