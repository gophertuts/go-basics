package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i32 int32
	var i64 int64
	var i int
	var s string
	var m map[int]complex128
	var es struct{}
	var ss struct {
		Name string
		Age  int32
	}

	fmt.Printf("int32: %d\n", unsafe.Sizeof(i32))
	fmt.Printf("int64: %d\n", unsafe.Sizeof(i64))
	fmt.Printf("int: %d\n", unsafe.Sizeof(i))
	fmt.Printf("string: %d\n", unsafe.Sizeof(s))
	fmt.Printf("map: %d\n", unsafe.Sizeof(m))
	fmt.Printf("empty struct: %d\n", unsafe.Sizeof(es))
	fmt.Printf("struct with fields: %d\n", unsafe.Sizeof(ss))
}
