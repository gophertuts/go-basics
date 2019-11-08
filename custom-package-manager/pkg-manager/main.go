package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	helpMsg := "\nUSAGE:" +
		"\npkg-manager help\t\t\t\t\t\t\t\t --- " +
		"displays this message\n" +
		"\npkg-manager init project-name\t\t\t\t\t\t\t --- " +
		"creates a new bare bones project\n" +
		"\npkg-manager get ./...\t\t\t\t\t\t\t\t --- " +
		"downloads all dependencies specified in 'go-pkg.json'\n" +
		"\npkg-manager get github.com/user1/repo@version github.com/user2/repo@version\t --- " +
		"downloads the mentioned list of dependencies using the git tag version and updates the go-pkg.json\n" +
		"\npkg-manager get github.com/user/repo@GIT_COMMIT_HASH\n" +
		"\npkg-manager get github.com/user/repo@GIT_COMMIT_HASH\t\t\t\t --- " +
		"downloads the mentioned list of dependencies using the git commit hash and updates the go-pkg.json\n"
	if len(args) > 0 && args[0] == "help" {
		fmt.Print(helpMsg)
		return
	}

	if len(args) < 2 {
		helpMsg = "Incorrect use of command\n" + helpMsg
		log.Fatal(helpMsg)
	}

	cmd := strings.TrimSpace(strings.ToLower(args[0]))
	cmdArgs := args[1:]
	m := NewManager([2]string{args[0], args[1]})

	switch {
	case cmd == "init" && len(cmdArgs) == 1:
		fmt.Println("=== creating a bare bones project ===")
		CreateProject(cmdArgs[0])
		UpdateGoPkg(m, map[string]string{})
	case cmd == "get" && cmdArgs[0] == "./...":
		fmt.Println("=== downloading all dependencies from go-pkg.json ===")
		Download(m.Dependencies)
	case cmd == "get" && cmdArgs[0] != "./...":
		deps := NewDependencies(cmdArgs)
		fmt.Println("=== downloading specified dependencies ===")
		Download(deps)
		fmt.Println("=== updating go-pkg.json ===")
		UpdateGoPkg(m, deps)
	}
}
