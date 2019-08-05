package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10

	ptr1 := unsafe.Pointer(&i)
	fmt.Println(*(*int)(ptr1))

	uptr1 := uintptr(i)
	uptr1 += 20

	ptr2 := unsafe.Pointer(&uptr1)
	fmt.Println(*(*int)(ptr2))
}
