package main

import (
	"fmt"
	"unsafe"
)

const (
	uInt8 int8 = 127
)

var (
	greatestUint32        = ^uint32(0)
	greatestUint64        = ^uint64(0)
	uInt64         uint64 = 12.0          // default conversion to int64
	uInt32                = int32(uInt64) // explicit conversion to int32
	f64                   = 12.5
	uInt64V2              = int64(f64) // explicit conversion to int64
	c                     = 'A'        // rune = int32
	r              rune   = 'B'
)

func main() {
	fmt.Println(greatestUint32)
	fmt.Println(greatestUint64)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(uInt64), uInt64)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(uInt32), uInt32)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(uInt64V2), uInt64V2)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(c), c)
	fmt.Printf("size: %d | type %T\n", unsafe.Sizeof(r), r)
	fmt.Println(uInt8)
	fmt.Println(unsafe.Sizeof(10)) // 32b/64b => 4/8B
}
