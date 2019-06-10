package main

import "fmt"

type person struct {
	Name string
	Age  int
}

const i = 10

var (
	s  = "some string"
	p1 = person{}
)

func main() {
	p2 := person{Name: "John", Age: 23}
	fmt.Printf("i: %T\n", i)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)
}
