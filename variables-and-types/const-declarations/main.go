package main

import (
	"fmt"
	"net/http"
)

// Computer measurement units
const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
	api = "api/v1"
)

// Week days
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("KB: %dB\n", KB)
	fmt.Printf("MB: %dB\n", MB)
	fmt.Printf("GB: %dB\n", GB)
	fmt.Printf("TB: %dB\n", TB)
	fmt.Printf("API prefix: %s\n", api)
	fmt.Printf("HTTP OK Status: %d\n", http.StatusOK)

	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
}
