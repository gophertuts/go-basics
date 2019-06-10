package main

import (
	"fmt"
	"runtime"
)

var (
	hostOS   string
	hostArch string
)

func main() {
	fmt.Println("runtime values:", runtime.GOOS, runtime.GOARCH)
	fmt.Println("build time values:", hostOS, hostArch)
}
