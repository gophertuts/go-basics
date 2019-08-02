package main

import (
	"fmt"
	"unsafe"
)

const (
	b byte = 1
)

func main() {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(false))
	fmt.Println(unsafe.Sizeof(1 > 10))
	fmt.Println(unsafe.Sizeof(!true))
	fmt.Println(unsafe.Sizeof(b))
}
