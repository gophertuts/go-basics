package pkg2

import (
	"fmt"

	// pkg3 side effects (init)
	_ "github.com/gophertuts/go-basics/go-program-anatomy/program-flow/pkg3"
)

const (
	c2 = 10
	// other consts
)

var (
	v2 = c2
	// other vars
)

func init() {
	fmt.Println("package pkg2: v2:", v2)
}
