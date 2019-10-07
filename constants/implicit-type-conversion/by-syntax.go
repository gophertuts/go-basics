package main

import "fmt"

type T complex128

const t T = 1

func main() {
	fmt.Printf("Value: %+v | Type: %T\n", 10, 10)
	fmt.Printf("Value: %+v | Type: %T\n", ^3, ^3)
	fmt.Printf("Value: %+v | Type: %T\n", 10 > 2, 10 > 2)
	fmt.Printf("Value: %+v | Type: %T\n", 1<<7, 1<<7)
	fmt.Printf("Value: %+v | Type: %T\n", 'A', 'A')
	fmt.Printf("Value: %+v | Type: %T\n", 11.2, 11.2)
	fmt.Printf("Value: %+v | Type: %T\n", 1e+15, 1e+15)
	fmt.Printf("Value: %+v | Type: %T\n", 1e-10, 1e-10)
	fmt.Printf("Value: %+v | Type: %T\n", .2, .2)
	fmt.Printf("Value: %+v | Type: %T\n", 12i, 12i)
	fmt.Printf("Value: %+v | Type: %T\n", "hello", "hello")
	fmt.Printf("Value: %+v | Type: %T\n", 1+2.5, 1+2.5)
	fmt.Printf("Value: %+v | Type: %T\n", 2/3, 2/3)
	fmt.Printf("Value: %+v | Type: %T\n", 10/2.0, 10/2.0)
	fmt.Printf("Value: %+v | Type: %T\n", 24./12, 24./12)
	fmt.Printf("Value: %+v | Type: %T\n", 1+13.2/t*5i-'D', 1+13.2/t*5i-'D')
}
