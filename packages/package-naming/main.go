package main

import (
	"github.com/gophertuts/go-basics/packages/package-naming/anotherWeirdPkg"
	"github.com/gophertuts/go-basics/packages/package-naming/collision"
	"github.com/gophertuts/go-basics/packages/package-naming/redundant"
	"github.com/gophertuts/go-basics/packages/package-naming/right"
	"github.com/gophertuts/go-basics/packages/package-naming/weird_long_package_name_with_underscores"
)

func main() {
	// GOOD
	redundant.NewMemory()
	// BAD
	// no need for the word 'Redundant', you already reference the package redundant
	redundant.NewRedundantMemory()

	// BAD
	weird_long_package_name_with_underscores.F()
	anotherWeirdPkg.W()

	// BAD - package and directory names are different
	// Don't be afraid of collision as long as the package is different
	wrong.New()
	collision.New()
}
