package pkg3

import "fmt"

func init() {
	fmt.Println("package pkg3: v3:", v3)
}

// not in expected order
var (
	v3 = c3
	// other vars
)

const (
	c3 = 100
	// other consts
)
