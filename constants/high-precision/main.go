package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Pi                = 3.14159265358979323846264338327950288419716939937510582097494459
	PiFloat64 float64 = Pi
	bb                = 22.0 / 7
)

func main() {
	//3.141592653589793
	//3.142857142857143
	fmt.Println(bb)
	f, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := strconv.FormatFloat(Pi, 'f', -1, 64)
	//s := fmt.Sprintf("%.64f\n%.64f", Pi, PiFloat64)

	_, err = f.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
