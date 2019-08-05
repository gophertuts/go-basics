package main

import (
	"fmt"
	"unsafe"
)

var (
	rs = []rune{'W', 'e', 'l', 'c', 'o', 'm', 'e', '好'}
	bs = []byte{'W', 'e', 'l', 'c', 'o', 'm', 'e', '!'}
)

func main() {
	s := "Hello World"
	// a string is just a slice of bytes or a sequence of uint8(s)
	fmt.Printf("%T\n", s[0])
	// check out the underlying string structure
	fmt.Println(unsafe.Sizeof(s)) // 64/128b

	// explicit conversion to string
	fmt.Println(string(rs))
	fmt.Println(string(bs))
}
