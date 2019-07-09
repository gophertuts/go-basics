package main

import (
	"fmt"
)

type T int64

const (
	t T = 1
	// keep the order in mind, the last from this list has top priority
	// integer, rune, floating-point, complex, custom type
	n =  2 + t * 'c' * 2.0
	shift = 'c' << 5
)

func main() {
	fmt.Printf("%T\n", n)
	fmt.Printf("%T", shift)
}
