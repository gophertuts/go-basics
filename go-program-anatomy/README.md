# Go program anatomy:

## Prerequisites

- Install Docker for your platform by checking
[Docker Installation](https://docs.docker.com/install/)

## How to test `go-build`

Create a Docker container to simulate the Linux environment
FYI: It's easier to run it inside Docker then spinning up a VM ;)

```bash
# download the repository with src code if you haven't already
go get github.com/gophertuts/go-basics

# cd into go-build directory on your machine
cd $GOPATH/src/github.com/gophertuts/go-basics/go-program-anatomy/go-build
```

```bash
# generate the binary 
go build
```

```bash
# spin up a docker container that simulates linux OS with Go runtime support
docker run --rm -it -v $(pwd):/go/src/github.com/gophertuts/go-basics/go-program-anatomy/go-build golang
```

```bash
# cd into $GOPATH and run the generated binary
cd /go/src/github.com/gophertuts/go-basics/go-program-anatomy/go-build
./go-build
```

## How to test `runtime-vs-buildtime`

```bash
# cd into go-build directory on your machine
cd $GOPATH/src/github.com/gophertuts/go-basics/go-program-anatomy/runtime-vs-buildtime
```

```bash
# generate the binary 
env GOOS=linux GOARCH=386 GOHOSTOS=linux GOHOSTARCH=386 \
go build -ldflags \
"-X main.hostOS=$(go env GOHOSTOS) -X main.hostArch=$(go env GOHOSTARCH)"
```

```bash
# spin up a docker container that simulates linux OS with Go runtime support
docker run --rm -it -v $(pwd):/go/src/github.com/gophertuts/go-basics/go-program-anatomy/runtime-vs-buildtime golang
```

```bash
# cd into $GOPATH and run the generated binary
cd /go/src/github.com/gophertuts/go-basics/go-program-anatomy/runtime-vs-buildtime
./runtime-vs-buildtime
```

#### Note:

`$GOOS` and `$GOARCH` are evaluated at compile time and saved as constants
inside `runtime.GOOS` and `runtime.GOARCH`

Also keep in mind that if `$GOOS` and `$GOARCH` are not set, the
default values will be taken from `$GOHOSTOS` and `$GOHOSTARCH`
which by the way `CANNOT` be changed at `compile/build time`


Play around with different `GOOS` & `GOARCH` values, here are some examples:

- `env GOOS=linux GOARCH=386 go build`
- `env GOOS=linux GOARCH=amd64 go build`
- `env GOOS=darwin GOARCH=386 go build`
- `env GOOS=darwin GOARCH=amd64 go build`

## Resources 💎

- [Go Build](https://golang.org/pkg/go/go-build/)
- [Go supported OS & Architectures](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
- [Go source code](https://golang.org/src/)

## FEEDBACK ⚗

[GopherTuts TypeForm](https://gophertuts.typeform.com/to/j2CJmC)

## COMMUNITY 🙌

[GopherTuts Discord](https://discord.gg/4sgecdh)

---

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>