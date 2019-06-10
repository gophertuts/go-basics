package main

import "fmt"

const (
	privateNumber = 10
	// PublicNumber is an exported symbol (constant)
	PublicNumber = 100
)

type privateType struct{}

// PublicType is an exported type
type PublicType struct{}

var (
	privateVar string
	// PublicVar is an exported symbol (variable)
	PublicVar string
)

func main() {
	fmt.Println("Exported and un exported symbols for a package example")
}
