package main

import (
	"fmt"
	"github.com/gophertuts/go-basics/constants/scope/p1"
)

const (
	Exported   = 10
	unExported = "Hello"
)

func main() {
	const (
		private = 10
		Private = "Hello"
	)

	fmt.Println("main: p1: Exported: ", p1.Exported)
	fmt.Println("main: Exported: ", Exported)
	fmt.Println("main: unExported: ", unExported)
	fmt.Println("main: Private: ", Private)
	fmt.Println("main: private: ", private)
}
