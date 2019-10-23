package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Manager struct {
	Dependencies map[string]string `json:"dependencies"`
}

func main() {
	fmt.Println("Reading dependencies from go-pkg.json")
	file, err := os.Open("go-pkg.json")
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := ioutil.ReadAll(file)

	var deps Manager
	err = json.Unmarshal(bs, &deps)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Downloading dependencies")
	Download(deps)
}
