package main

import (
	"fmt"

	_ "github.com/gophertuts/go-basics/go-program-anatomy/program-flow/pkg1"
)

const (
	mc = "const"
)

var (
	mv = mc + ":var"
)

func init() {
	mv += ":init"
}

func main() {
	mv += ":main"
	fmt.Println("package main: mv:", mv)
}
