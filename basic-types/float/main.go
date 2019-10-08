package main

import (
	"fmt"
	"unsafe"
)

var (
	f32   float32 = 5    // automatic conversion to float32
	f64   float64 = 13   // automatic conversion to float64
	f64_2         = 13.2 // automatic conversion to float64
)

func main() {
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(f32), f32)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(f64), f64)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(f64_2), f64_2)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(.2), .2)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(12.5), 12.5)
	fmt.Printf("value: %[1]v | type: %[1]T\n", 2*3.0) // conversion to float64
	fmt.Printf("value: %[1]v | type: %[1]T\n", 14/3)  // no conversion is done
}
