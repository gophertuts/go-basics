package main

import "fmt"

const s = "Hello"

var bs = []rune{'H', 'e', 'l', 'l', 'o'}

func main() {
	// cannot assign to s[0]
	// s[0] = 'M'
	fmt.Println(s)

	bs[0] = 'M'
	fmt.Println(string(bs))
}
