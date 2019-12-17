package main

import "github.com/steevehook/some-test/pkg"

func main() {
	// this won't even compile
	// because of import path checking
	pkg.New()
}
