package main

import "os"

func main() {
	args := os.Args[1:]
	manager := ReadGoPkg()
	CommandSwitch(args, manager)
}
