package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) IsAdult() bool {
	return p.Age > 18
}

func main() {
	p := &Person{
		Name: "Mukesh",
		Age:  28,
	}

	fmt.Println(p.GetName())
	fmt.Println(p.IsAdult())

}
