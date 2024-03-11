package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) speak() {
	fmt.Println("Person speaks")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := Person{
		Name: "Mukesh",
		Age:  28,
	}

	// saySomething(p) // cannot pass Person
	saySomething(&p) // Can only pass pointer to Person
	p.speak()
}
