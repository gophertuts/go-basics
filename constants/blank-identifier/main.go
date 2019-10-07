package main

import (
	// unused import
	_ "fmt"
	// package side effects (like init)
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// unused const identifier
const (
	_ = 10
)

func main() {
	// unused (ignored) error value
	f, _ := os.Open("filename.txt")
	log.Print(f.Name())
}
