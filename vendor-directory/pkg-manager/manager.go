package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Manager struct {
	Dependencies map[string]string `json:"dependencies"`
}

func NewDependencies(packageList []string) map[string]string {
	deps := map[string]string{}
	for _, dep := range packageList {
		ds := strings.Split(dep, "@")
		pkg := ds[0]
		var version string
		if len(ds) < 2 {
			version = "master"
		} else {
			version = ds[1]
		}
		deps[pkg] = version
	}
	return deps
}

func UpdateGoPkg(m Manager, newDeps map[string]string) {
	for dep, ver := range newDeps {
		m.Dependencies[dep] = ver
	}
	// JSON marshall the manager
	bs, _ := json.Marshal(m)
	bs = append(bs, '\n')

	// write the new result to the file
	err := ioutil.WriteFile("go-pkg.json", prettyPrint(bs), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadGoPkg() Manager {
	// read go-pkg.json file
	fmt.Println("Reading dependencies from go-pkg.json")
	file, err := os.Open("go-pkg.json")
	if err != nil {
		log.Fatalf("go-pkg.json not found, create one\n%+v", err)
	}
	bs, _ := ioutil.ReadAll(file)

	// unmarshal the json into the struct
	var m Manager
	err = json.Unmarshal(bs, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func prettyPrint(b []byte) []byte {
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "  ")
	return out.Bytes()
}
