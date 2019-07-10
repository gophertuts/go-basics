package main

import "fmt"

const (
	// typed
	b    bool       = true
	i    int        = -64 - 32
	i8   int8       = -8
	i16  int16      = -16
	i32  int32      = -32
	r    rune       = 'c' // alias of in32
	i64  int64      = -64
	u    uint       = 1e2
	u8   uint8      = 8
	bb   byte       = 20 // alias of uint8
	u16  uint16     = 16
	u32  uint32     = 32
	u64  uint64     = 64
	uptr uintptr    = '\u0001'
	f32  float32    = 3.2
	f64  float64    = 6.4
	c64  complex64  = 6 + 4i
	c128 complex128 = 12 + 8i
	str  string     = "Hello"

	// untyped
	untypedB       = false
	untypedInt     = 10
	untypedIntIota = iota // alias of untyped int
	untypedRune    = '汉'
	untypedFloat   = 2.0
	untypedComplex = 2 + 0i
	untypedString  = "World"

	// const expressions
	exp1 = 4.5 + 1 + 2i
	exp2 = 12 + 3.5
	exp3 = 'H' + 'e' + 'l' + 'l' + 'o' + ' ' + 'W' + 'o' + 'r' + 'l' + 'd'
	exp4 = str + " " + untypedString
)

func main() {
	fmt.Printf("exp1: Value: %v Type: %T\n", exp1, exp1)
	fmt.Printf("exp2: Value: %v Type: %T\n", exp2, exp2)
	fmt.Printf("exp3: Value: %v Type: %T\n", exp3, exp3)
	fmt.Printf("exp4: Value: %v Type: %T\n", exp4, exp4)
}
