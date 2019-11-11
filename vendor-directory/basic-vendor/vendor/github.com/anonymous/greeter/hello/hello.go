package hello

import (
	"fmt"

	"github.com/anonymous/greeter/chinese"
	"github.com/gophertuts/go-basics/vendor-directory/basic-vendor/en"
	"github.com/gophertuts/go-basics/vendor-directory/basic-vendor/ru"
)

func Greet() {
	fmt.Println("Hello from anonymous")
	chinese.HelloChinese()
	en.HelloEnglish()
	ru.HelloRussian()
}
