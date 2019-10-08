package main

import (
	"fmt"
	"unsafe"
)

const (
	b1 byte = 1
	b3 byte = 255 // if greater than 255: const will overflow byte
)

var (
	b2 byte = 'A'
)

func main() {
	fmt.Printf("%T\n", b1)
	fmt.Println(unsafe.Sizeof(b1))
	fmt.Println("uint value", b2, "string value", string(b2))
	fmt.Println(b3)
	// automatic conversion to 'uint8'
	fmt.Printf("value of: A + B = %[1]v | type: %[1]T\n", 'A'+b2)
	// no conversion is done, falling back to 'int32' aka 'rune'
	fmt.Printf("value of: A + B = %[1]v | type: %[1]T\n", 'A'+'B')
}
