package main

import "fmt"

const (
	// untyped bool - bool
	b = false

	// untyped int - int
	i = 12

	// untyped int - int
	j = iota

	// untyped rune - int32
	c = 'N'

	// untyped float - float64
	f = 3.5
	e = 1e+200

	// untyped complex - complex128
	cx = 3 + 4i

	// untyped string - string
	s = "Hello World"
)

func main() {
	fmt.Printf("b: %T\n", b)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("j: %T\n", j)
	fmt.Printf("c: %T\n", c)
	fmt.Printf("f: %T\n", f)
	fmt.Printf("e: %T\n", e)
	fmt.Printf("cx: %T\n", cx)
	fmt.Printf("s: %T\n", s)
}
