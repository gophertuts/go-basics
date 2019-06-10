package main

import "fmt"

func main() {
	i := 10
	n, m := 20, 30
	// i := 20 // error: no new variables on left side of :=
	i, j := 40, 50

	fmt.Println("Value of i is ", i)
	fmt.Println("Value of j is ", j)
	fmt.Printf("Value of n is %d and m is %d\n", n, m)
}
