# Basic Types in Go

## Overall types discussed in this tutorial:

- bool
- int
- int8
- int16
- int32/rune
- int64
- int64
- uint
- uint8/byte
- uint16
- uint32
- uint64
- uint64
- uintptr
- float32
- float64
- complex64
- complex128
- string

Note: `int`, `uint`, `uintptr`, `string` sizes are platform dependent
In case you want to simulate x32 or x64 platform architecture
follow the instructions below

## Simulate x32 / x64 bit architecture

```bash
# cd into the project i.e. int examples
cd $GOPATH/src/github.com/gophertuts/go-basics/basic-types/int

# 32bit
env GOARCH=386 go run main.go

# 64bit
env GOARCH=amd64 go run main.go

# $GOARCH => $GOHOSTARCH will be taken
go run main.go

# to check you host architecture run
go env | grep GOHOSTARCH
```

## Resources

- [String implementation](https://github.com/golang/go/blob/master/src/runtime/string.go)
- [Strings - Go101](https://go101.org/article/string.html)
- [Unsafe package](https://golang.org/pkg/unsafe/)
- [Unsafe - Go101](https://go101.org/article/unsafe.html)
- [UintPtr vs Unsafe pointer](https://utcc.utoronto.ca/~cks/space/blog/programming/GoUintptrVsUnsafePointer)
- [Complex numbers](https://en.wikipedia.org/wiki/Complex_number)

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>