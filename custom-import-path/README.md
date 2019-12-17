# Custom **import** path in Go

## Overview

The purpose of this tutorial is to teach you what are custom import
paths in Go and how can you make use of this feature to your advantage.

In this tutorial you will learn how to set up a custom import path
static Go server, which tells us the client which cloud provider
to use for our import and we'll have a small program
which uses that package.

The main topics we're gonna cover are:

- `Custom import paths`
- `Import path checking`

## Medium version 📖

[Custom import paths in Go](https://medium.com/@gophertuts/packages-in-go-df5438123548)

### How to run:

1. Custom `go get` **import** path

```bash
# cd into the example
cd $GOPATH/src/github.com/gophertuts/go-basics/custom-import-path/custom-path

# fetch the package from custom URL
go get -insecure -u 0.0.0.0/user/pkg

# spin an HTTP static server
go run server.go
```

#### Run GitLab Docker server:

***Disclaimer***: GitLab says it's not guaranteed to work on Windows,
so if you happen to use Windows. You can try this out, but not
sure it's going to reproduce or work properly.

You can though try it on Linux, i.e. Ubuntu

```bash
# run the docker container and expose port 443 and 8000
# don't forget to map the volumes correctly as shown below
# we use port 8000 to not collapse with port 80 which we need by server.go
docker run --detach \
  --hostname gitlab.example.com \
  --publish 443:443 --publish 8000:80 \
  --restart always \
  --volume ~/Desktop/gitlab/config:/etc/gitlab \
  --volume ~/Desktop/gitlab/log:/var/log/gitlab \
  --volume ~/Desktop/gitlab/opt:/var/opt/gitlab \
  gitlab/gitlab-ce:latest
  
# display real time logs for GitLab docker container
docker logs --follow [CONTAINER_ID]

# WAIT TILL IT FINISHES

# open up localhost:8000

# set up a new password, preferably something which uses
# Uppercase, letters, numbers, alphanumeric characters
# otherwise GitLab won't allow you to go further
# something like: Pas$word123456

# login with the credentials:
# username: root
# password: [PASSWORD_YOU_CREATED]

# create a repository from the admin panel and name it `some-test`

# clone the created repository on your local machine
git clone http://localhost:8000/root/some-test.git

# add the go pkg shown bellow and save it inside pkg.go

git add -A

git commit -m "initial commit"

git push

# input your username: root
# input your password: [PASSWORD_YOU_CREATED]

# DONE
```

###### `pkg.go`
```go
package pkg // import "0.0.0.0/user/pkg"

import "fmt"

func New() {
	fmt.Println("Welcome to pkg! (GitLab - Docker)")
}
```

#### More info

***Note***: import path checking is disabled when using
`Go modules` or `vendor` directory

```bash
# displays useful information about import paths in Go
go help importpath

# displays useful information about Go specific env variables
# i.e. GO111MODULE
go help environment
```

## Resources 💎

- [`go get` custom import path](https://jve.linuxwall.info/blog/index.php?post/2015/08/26/Hosting_Go_code_on_Github_with_custom_import_path)
- [Remote import paths](https://golang.org/cmd/go/#hdr-Remote_import_paths)
- [Custom import path checking](https://docs.google.com/document/d/1jVFkZTcYbNLaTxXD9OcGfn7vYv5hWtPx9--lTx1gPMs/edit)
- [GitLab Docker](https://docs.gitlab.com/omnibus/docker/)

## FEEDBACK ⚗

[GopherTuts TypeForm](http://feedback.gophertuts.com)

## COMMUNITY 🙌

[GopherTuts Discord](https://discord.gg/4sgecdh)

---

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>
