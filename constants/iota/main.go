package main

import "fmt"

// week days
const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// byte multiples
const (
	// first value is ignored (we do not want 0)
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println("Week days:")
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println("Byte multiples:")
	fmt.Println(KB, MB, GB)
}
