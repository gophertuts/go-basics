package main

import "fmt"

const (
	u uint = 1e9
	i int  = 1
)

func main() {
	// invalid operation: i + u (mismatched types int and uint)
	//fmt.Println("Result", i+u)
	fmt.Println("Result1", uint(i) + u)
	fmt.Println("Result2", int(u) + i)
}
