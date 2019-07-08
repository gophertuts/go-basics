# Constants in Go

## Types of const(s)

- Typed
- Untyped (Kind)

## Const declaration examples

```go
package main
const (
	// bool
	t = true
	f bool = false
	
	// numeric
	num1 = 10
	num2 int64 = 20
	num3 = 15 + 3i
	
	// overflows
	huge = 10 << 100
	
	// rune (char)
	r1 = 'A'
	r2 = '\u212A'
	r3 = '你'
	
	// string
	hello_world = "Hello World"
	_underscore = "Lodash"
	
	// ignored
	_ = 125 + 13i
	
	// visibility
	Exported = 10
	unexported = 0
)
```

## Iota examples

```go
package main

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
	_ = iota
	KB = 1 << (iota * 10)
	MB
	GB
)
```

## C Background

## Resources

- [Constants](https://blog.golang.org/constants)
- [Basic Kind](https://golang.org/pkg/go/types/#Basic.Kind)
- [Iota](https://github.com/golang/go/wiki/Iota)
- [Identifiers](https://golang.org/ref/spec#Identifiers)
- [Keywords](https://golang.org/ref/spec#Keywords)
- [Numeric types](https://golang.org/ref/spec#Numeric_types)
- [Introduction to Numeric Constants](https://www.ardanlabs.com/blog/2014/04/introduction-to-numeric-constants-in-go.html)
- [A Tour of Go - Basics](https://tour.golang.org/basics/1)
