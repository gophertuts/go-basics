package main

import (
	"fmt"
	"unsafe"
)

var (
	c64    complex64  = 12
	c128   complex128 = 128
	c128_2            = 3i
	c128_3            = +0i
	c64_2  complex64  = 3 + 15
)

func main() {
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(c64), c64)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(c128), c128)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(c128_2), c128_2)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(c128_3), c128_3)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(c64_2), c64_2)
	fmt.Println(unsafe.Sizeof(1 + 2i)) // 128b => 16B
	fmt.Println(11 + 25.3*4 + 4i)      // real part - 112.2 | imaginary part - 4i
	fmt.Println(4i)                    // real part - 0 | imaginary part - 4i
}
