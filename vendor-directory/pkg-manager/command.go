package main

import (
	"fmt"
	"log"
	"strings"
)

func CommandSwitch(args []string, m Manager) {
	helpMsg := "\ninvalid use of 'pkg-manager' command\n" +
		"\npkg-manager get ./...\t\t\t\t\t\t\t\t --- " +
		"downloads all dependencies specified in 'go-pkg.json'\n" +
		"pkg-manager get github.com/user1/repo@version github.com/user2/repo@version\t --- " +
		"downloads the mentioned list of dependencies and updates the go-pkg.json\n" +
		"pkg-manager get github.com/user/repo@GIT_COMMIT_HASH\n"
	if len(args) < 2 {
		log.Fatal(helpMsg)
	}
	cmd := strings.TrimSpace(strings.ToLower(args[0]))
	packageList := args[1:]

	switch {
	case cmd == "get" && packageList[0] == "./...":
		fmt.Println("Downloading all dependencies from go-pkg.json")
		Download(m.Dependencies)
	case cmd == "get" && packageList[0] != "./...":
		fmt.Println("Updating go-pkg.json")
		deps := NewDependencies(packageList)
		UpdateGoPkg(m, deps)
		fmt.Println("Downloading specified dependencies")
		Download(deps)
	default:
		log.Fatal(helpMsg)
	}
}
