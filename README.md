# gear

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/joseph-beck/gear/checks.yml?label=checks)
![GitHub Release](https://img.shields.io/github/v/release/joseph-beck/gear)
![GitHub License](https://img.shields.io/github/license/joseph-beck/gear)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/joseph-beck/gear?label=go)

```txt                     
    ▄▄             ▄    
 ▄████ ▄█▀█▄ ▄▀▀█▄ ████▄
 ██ ██ ██▄█▀ ▄█▀██ ██   
▄▀████▄▀█▄▄▄▄▀█▄██▄█▀   
    ██                  
  ▀▀▀                   
```

## about

a peg parsing library written in go using packrat parsing to parse left recursive expressions!

## usage

to install `gear`

```bash
# install gear in your project
$ go get -u github.com/joseph-beck/gear
```

```go
// import with
import (
    "github.com/joseph-beck/gear/pkg/gear"
)
```

for examples see [examples](/example/)

## getting started

```bash
# 1. make sure go is installed
# go to https://go.dev/doc/install for to install

# 2. install go packages
make install

# 3. run
make cli

# etc. make commands
# view all other commands and their use
make help
```
