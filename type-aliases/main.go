package main

import "fmt"

// alias of int
type aType = int

// alias of int
type bType int

func (b bType) String() string {
	return "some string"
}

// alias of bType
type cType = bType

func main() {
	var x int
	var i aType
	i = 20
	i = x
	fmt.Println(i)

	b := bType(1)
	fmt.Println(b)

	c := cType(1)
	fmt.Println(c)

	// alias of int32
	var r rune // type rune = int32
	r = 65
	fmt.Println(string(r))
}
