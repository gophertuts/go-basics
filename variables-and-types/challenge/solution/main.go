package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var n int
	fileName := flag.String("filename", "", "The name of the file to read contents from")
	flag.Parse()

	if strings.Trim(*fileName, "") == "" {
		log.Fatal("filename flag was not provided")
	}

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n++
	}

	fmt.Printf("The file has %d lines\n", n)
}
