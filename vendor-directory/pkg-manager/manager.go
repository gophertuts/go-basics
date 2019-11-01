package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Manager struct {
	WorkDir      string            `json:"-"`
	Dependencies map[string]string `json:"dependencies"`
}

func NewManager(args [2]string) Manager {
	if strings.TrimSpace(strings.ToLower(args[0])) == "init" {
		return Manager{
			WorkDir:      args[1],
			Dependencies: map[string]string{},
		}
	}
	var m Manager
	err := json.Unmarshal(readGoPkg(), &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
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

func CreateProject(projectName string) {
	fmt.Println("=== creating project directory ===")
	err := os.Mkdir(path.Join(".", projectName), 0755)
	if err != nil {
		log.Fatalf("could not create project directory\n%+v", err)
	}
	fmt.Println("=== creating empty go-pkg.json file ===")
	_, err = os.Create(path.Join(".", projectName, "go-pkg.json"))
	if err != nil {
		log.Fatalf("could not create go-pkg.json file\n%+v", err)
	}
}

func UpdateGoPkg(m Manager, newDeps map[string]string) {
	for dep, ver := range newDeps {
		m.Dependencies[dep] = ver
	}
	// JSON marshall the manager
	bs, _ := json.Marshal(m)
	bs = append(bs, '\n')

	// write the new result to the file
	err := ioutil.WriteFile(path.Join(".", m.WorkDir, "go-pkg.json"), prettyPrint(bs), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func readGoPkg() []byte {
	file, err := os.Open("go-pkg.json")
	if err != nil {
		log.Fatalf("go-pkg.json not found, create one\n%+v", err)
	}
	bs, _ := ioutil.ReadAll(file)
	return bs
}

func prettyPrint(b []byte) []byte {
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "  ")
	return out.Bytes()
}
