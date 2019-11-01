package main

import (
	"fmt"

	"github.com/anonymous/greeter/hello"
	"github.com/anonymous/greeter/world"
	"github.com/julienschmidt/httprouter"
	johnGreeter "gitlab.com/john/strings/greeter"
	steveGreeter "gitlab.com/steve/strings/greeter"
)

func main() {
	hello.Greet()
	world.Greet()
	johnGreeter.Greet()
	steveGreeter.Greet()
	router := httprouter.New()
	fmt.Println("I am just a number:", router)
}
