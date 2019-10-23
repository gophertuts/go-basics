package main

import (
	"github.com/anonymous/greeter/hello"
	"github.com/anonymous/greeter/world"

	johnGreeter "gitlab.com/john/strings/greeter"
	steveGreeter "gitlab.com/steve/strings/greeter"
)

func main() {
	hello.Greet()
	world.Greet()
	johnGreeter.Greet()
	steveGreeter.Greet()
}
