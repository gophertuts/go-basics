package main

import "fmt"

type custom struct {
	name string
	age  int
}

var (
	i, j       int
	f          float64
	c          complex128
	r          rune
	s          string
	m          map[string]bool // nil map :(
	sentences  []string
	people     struct{}
	anything   interface{}
	fn         func(bool) error
	customType custom
)

func main() {
	if m == nil {
		fmt.Println("We got a nil map")
	}

	fmt.Printf("i: %+v\n", i)                 // 0
	fmt.Printf("j: %+v\n", j)                 // 0
	fmt.Printf("f: %+v\n", f)                 // 0
	fmt.Printf("c: %+v\n", c)                 // (0+0i)
	fmt.Printf("r: %+v\n", r)                 // 0
	fmt.Printf("s: %+v\n", s)                 // ""
	fmt.Printf("m: %+v\n", m)                 // map[] // nil map
	fmt.Printf("sentences: %+v\n", sentences) // [] // empty slice
	fmt.Printf("people: %+v\n", people)       // {} // empty struct
	fmt.Printf("anything: %+v\n", anything)   // nil
	fmt.Printf("fn: %+v\n", fn)               // nil
	fmt.Printf("custom: %+v\n", customType)   // {name: "", age: 0}
}
