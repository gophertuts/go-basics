package main

import "fmt"

const interpreted = "This text is interpreted\n\n \tsome tabulated text\nand a special char '\u0028'"

func main() {
	fmt.Println(interpreted)
}
