package pkg1

import (
	"fmt"

	// pkg2 side effects (init)
	_ "github.com/gophertuts/go-basics/go-program-anatomy/program-flow/pkg2"
)

const (
	c1 = 1
	// other consts
)

var (
	v1 = c1
	// other vars
)

func init() {
	fmt.Println("package pkg1: v1:", v1)
}
