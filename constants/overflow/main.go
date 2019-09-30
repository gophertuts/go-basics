package main

import (
	"fmt"
	"math"
)

const Huge = 1.2676506002282294e+30

func main() {
	// int32
	fmt.Println(^int32(math.Pow(2, 31))) // 2147483647
	fmt.Println(int32(math.Pow(2, 31)))  // -2147483648

	// int64
	fmt.Println(^int64(math.Pow(2, 63))) // -9223372036854775808
	fmt.Println(int64(math.Pow(2, 63)))  // 9223372036854775807

	// constants/overflow/main.go:11:13: constant 1267650600228229400000000000000 overflows int
	fmt.Println(int(Huge))
}
