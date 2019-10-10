package main

import (
	"fmt"

	"github.com/gophertuts/go-basics/packages/exported-vs-unexported/pkg"
)

func main() {
	//fmt.Println(pkg.unExported)
	fmt.Println(pkg.Exported)
	fmt.Println(pkg.Public())
	//fmt.Println(pkg.private())
}
