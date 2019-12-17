# Import path checking

## Overview

This program proves how `import path checking` works
when enabled from a certain package.

***Note***: this is a feature designed mainly for
custom import paths which use `custom domains (hosts)`
instead of fixed hosting services
such as `GitHub`, `BitBucket`, `GitLab` and others.

```bash
# NOT GOING TO WORK (COMPILE)
# but the package will be downloaded inside
# $GOPATH/src/github.com/steevehook/some-test
go get github.com/steevehook/some-test

# Will fail, because of wrong import path
go run main.go
```
