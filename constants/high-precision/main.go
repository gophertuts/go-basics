package main

import (
	"fmt"
)

const (
	// constants are only available during compile time
	Pi                      = 3.14159265358979323846264338327950288419716939937510582097494459
	HugeMathematicallyExact = 36893488147419103230
	MaxUint                 = 18446744073709551615
	//MaxUint = ^uint(0)
)

func main() {
	fmt.Println(Pi)                     //3.141592653589793
	fmt.Println(3.141592653589793 / Pi) // 0.9999999999999999

	// constant 36893488147419103230 overflows int
	//fmt.Println(uint(HugeMathematicallyExact))
	fmt.Println(HugeMathematicallyExact / MaxUint) // 2
}
