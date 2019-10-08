package p1

import (
	"fmt"
	"github.com/gophertuts/go-basics/constants/scope/p2"
)

const (
	Exported   = "exported"
	unExported = 10
)

func init() {
	fmt.Println("p1: unExported:", unExported)
	fmt.Println("p1: p2: Exported:", p2.Exported)
}
