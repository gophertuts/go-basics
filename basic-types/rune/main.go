package main

import (
	"fmt"
	"unsafe"
)

var (
	r1      = 65
	r2      = 'Z'
	r3 rune = 128
)

func main() {
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof('A'+'B'), 'A'+'B')
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(r3), r3)
	fmt.Printf("some char range regex: [%s-%s]\n", string(r1), string(r2))
	fmt.Println("Your order total is of 100", string(r3))
}
