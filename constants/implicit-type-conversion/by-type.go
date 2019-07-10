package main

import "fmt"

var (
	i int        = 13.0
	f float64    = 5
	c complex128 = 12.3 + 5
)

func main() {
	fmt.Printf("Value: %+v | Type: %T\n", i, i)
	fmt.Printf("Value: %+v | Type: %T\n", f, f)
	fmt.Printf("Value: %+v | Type: %T\n", c, c)
}
