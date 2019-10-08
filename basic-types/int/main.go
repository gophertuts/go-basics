package main

import (
	"fmt"
	"unsafe"
)

const (
	i8 int8 = -128
)

var (
	i64   int64 = 12.0       // default conversion to int64
	i32         = int32(i64) // explicit conversion to int32
	f64         = 12.5
	i64_2       = int64(f64) // explicit conversion to int64
	c           = 'A'        // rune = int32
	r     rune  = 'B'
)

func main() {
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(i64), i64)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(c), c)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(r), r)
	fmt.Println(i8)
	fmt.Println(unsafe.Sizeof(10)) // 32b/64b => 4/8B
}
