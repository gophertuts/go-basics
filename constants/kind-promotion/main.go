package main

import (
	"fmt"
)

type T complex128

const (
	t T = 1
	// keep the order in mind, the last from this list has the top priority in the end result type
	// integer, rune, floating-point, complex, string, custom type
	n     = 2 + t*'c'*2.0 + 35i
	shift = 'c' << 5
)

func main() {
	fmt.Printf("%T\n", n)
	fmt.Printf("%T", shift)
}
