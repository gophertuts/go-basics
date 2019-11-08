package main

import (
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

const vendor = "vendor"

func Download(deps map[string]string) {
	for dependency, version := range deps {
		_, err := os.Stat(vendor)
		if os.IsNotExist(err) {
			createVendor(dependency)
		} else {
			deleteVendor(dependency)
			createVendor(dependency)
		}
		repo, err := git.PlainClone(
			path.Join(".", vendor, dependency),
			false, &git.CloneOptions{
				URL:      "https://" + dependency,
				Progress: os.Stdout,
			})
		if err != nil {
			log.Fatalf("%+v: %s", err, dependency)
		}
		checkout(repo, version)
	}
}

func createVendor(dependency string) {
	err := os.MkdirAll(
		path.Join(".", vendor, dependency),
		0755,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteVendor(dependency string) {
	err := os.RemoveAll(path.Join(".", vendor, dependency))
	if err != nil {
		log.Fatal(err)
	}
}

func versionToHash(repo *git.Repository, version string) string {
	if version == "master" {
		head, _ := repo.Head()
		return head.String()
	}
	tags, _ := repo.Tags()
	for {
		tag, err := tags.Next()
		if err != nil {
			break
		}
		if strings.Contains(tag.Name().String(), version) {
			return tag.Hash().String()
		}
	}
	commits, _ := repo.CommitObjects()
	for {
		commit, err := commits.Next()
		if err != nil {
			break
		}
		if strings.Contains(commit.Hash.String(), version) {
			return commit.Hash.String()
		}
	}
	return ""
}

func checkout(repo *git.Repository, version string) {
	workTree, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	hash := versionToHash(repo, version)
	if hash == "" {
		cfg, _ := repo.Config()
		log.Fatalf(
			"no such version: '%s' for origin remote: '%s'",
			version,
			cfg.Remotes["origin"].URLs[0],
		)
	}
	if err = workTree.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(versionToHash(repo, version)),
	}); err != nil {
		log.Fatal(err)
	}
}
