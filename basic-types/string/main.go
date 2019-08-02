package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Hello World"
	fmt.Printf("%T\n", s[0])
	fmt.Println(unsafe.Sizeof(s))
}
