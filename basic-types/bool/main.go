package main

import (
	"fmt"
	"unsafe"
)

const num = 10

func main() {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(false))
	fmt.Println(unsafe.Sizeof(1 > 10))
	fmt.Println(unsafe.Sizeof(!true))
	if num > 3 {
		fmt.Println(`"num" is greater than 3`)
	}
}
