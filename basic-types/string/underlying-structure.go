package main

import (
	"fmt"
	"unsafe"
)

// underlying string type
type _string struct {
	// 4B/8B => 32/64b - depending on GOARCH
	elements *byte // underlying bytes
	// 4B/8B => 32/64b - depending on GOARCH
	len int // number of bytes
}

func main() {
	var b byte
	var pb *byte
	s := "Hello World"
	fmt.Println(unsafe.Sizeof(_string{})) // 8/16B => 128b
	fmt.Println(unsafe.Sizeof(pb)) // 4/8B => 64b
	fmt.Println(unsafe.Sizeof(b))  // 4/8B => 64b
	fmt.Println(unsafe.Sizeof(s))  // 8/16B => 128b
}
