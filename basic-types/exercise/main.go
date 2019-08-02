package main

import "fmt"

func main() {
	s := []byte("你 好")
	if len(s) == 3 {
		fmt.Println("String has 3 Bytes")
	} else {
		fmt.Println("String has 7 Bytes")
	}
	fmt.Println(len(s))
}
