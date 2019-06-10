package main

import "fmt"

var n1 int
var n2 float32 = 2.5
var n3 = 22
var people = []string{"Andrew", "John", "Lucy"}
var b *bool

func main() {
	f := false
	b = &f

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(*b)
	fmt.Println(people)
}
