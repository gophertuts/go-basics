# Constants in Go

## Types of const(s)

- Typed
- Untyped (Kind)

## Facts about constants

- Cannot be reassigned or redeclared
- Can only hold scalar values
- Can hold large high precision numbers
- Can be created from expression of other constants
- Can be untyped (Kind)
- Constants only **exist** during **compilation time**

## C background

#### How to run the C program

```bash
# cd into the project directory
cd $GOPATH/src/github.com/gophertuts/go-basics/constants/c-vs-go

# compile the program
gcc main.c -o exec

# run the program
./exec
```

## Resources 💎

- [Constants](https://blog.golang.org/constants)
- [Basic Kind](https://golang.org/pkg/go/types/#Basic.Kind)
- [Iota](https://github.com/golang/go/wiki/Iota)
- [Identifiers](https://golang.org/ref/spec#Identifiers)
- [Keywords](https://golang.org/ref/spec#Keywords)
- [Numeric types](https://golang.org/ref/spec#Numeric_types)
- [Representability](https://golang.org/ref/spec#Representability)
- [Constant expressions](https://golang.org/ref/spec#Constant_expressions)
- [Introduction to Numeric Constants](https://www.ardanlabs.com/blog/2014/04/introduction-to-numeric-constants-in-go.html)
- [A Tour of Go - Basics](https://tour.golang.org/basics/1)

## FEEDBACK ⚗

[GopherTuts TypeForm](http://feedback.gophertuts.com)

## COMMUNITY 🙌

[GopherTuts Discord](https://discord.gg/4sgecdh)

---

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>
