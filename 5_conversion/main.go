package main

import "fmt"

type custom int

var x custom

type newOne custom

var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
