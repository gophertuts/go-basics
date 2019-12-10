package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"

	_ "github.com/gophertuts/go-basics/vendor-directory/pkg1"
	_ "github.com/gophertuts/go-basics/vendor-directory/pkg2"
	_ "github.com/gophertuts/go-basics/vendor-directory/pkg3"
)

func main() {
	fmt.Println("some packages imported")
	router := httprouter.New()
	fmt.Println("router is just a number:", router)
}
