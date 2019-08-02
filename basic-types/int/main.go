package main

import (
	"fmt"
	"unsafe"
)

const (
	i8 int8 = -128
)

func main() {
	fmt.Println(i8)
	fmt.Println(unsafe.Sizeof(10))
}
