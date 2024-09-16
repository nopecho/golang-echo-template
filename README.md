# Go Template
<img alt="Go" height="" src="https://go.dev/doc/gopher/gopherbw.png" width="300"/>

## Introduction

>[한국어](README.kr.md)

This is a template project for building Go server applications using the [echo](https://echo.labstack.com/) framework.

It includes:
* [Version Management](#version-management)
* [Package Layout](#package-layout)
* [Various Scripts](#getting-started)
* [Profiling](#profiling)
* Sample Code

## Dependency

The following libraries are used in this template:

> The Go community generally recommends using the standard library whenever possible.
>
> However, for ease of development, we include a few essential libraries.
>
> _~~It's better not to reinvent the wheel.~~_

* Logging
    * [zerolog](https://github.com/rs/zerolog)
* Testing
    * [testify](https://github.com/stretchr/testify)
    * [testcontainers](https://golang.testcontainers.org/)
* HTTP Server
    * [echo](https://echo.labstack.com/)
* Database
    * [gorm](https://gorm.io/)

## Version Management

* Go version: `1.23.0`

The [asdf](https://asdf-vm.com/guide/introduction.html) tool is used for version management.

If `asdf` is not installed, refer to the [section below](#asdf) for installation instructions.

### Version Setting
```shell
make asdf
```

### asdf
>[asdf installation guide](https://asdf-vm.com/guide/getting-started.html)
>
>[Install Golang plugin](https://github.com/asdf-community/asdf-golang)

1. Install asdf
```shell
brew install asdf
```
2. Install plugin
```shell
asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
```
3. Set `GOROOT`
```shell
. ~/.asdf/plugins/golang/set-env.zsh
```

## Package Layout
```
├─ cmd
│  ├── web
│  ├── worker
│  └── etc..
├─ internal
│  ├── app
│  │   ├── api
│  │   ├── domain
│  │   ├── infra
│  │   └── svc
│  └── utils
├─ examples
└─ local
```
The package structure follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

* `cmd`: The main entry points for the application. This layer is responsible for starting the server and connecting all layers of the application.

* `internal`: Contains the application's internal packages.
    * `app`: This package holds the core logic of the application.
    * `utils`: These are packages that can be used across various parts of the application.

* `local`: Contains scripts and configuration files for local development.

## Getting Started

### Run

```shell
make run module=<module name>
```

### Build

```shell
 make build module=<module name>
```

### Docker Build

```shell
 make docker module=<module name> tag=<tag name>
```

### Test

```shell
 make test
```

## Profiling
This guide explains how to execute pprof for application profiling.

### Prerequisite
#### 1. Install `pprof`
```shell
go install github.com/google/pprof@latest
```

#### 2. Install `graphviz`
```shell
brew install graphviz
```

#### 3. Enable pprof in the application
```go
package main
import _ "net/http/pprof"

func main() {
	// pprof endpoint
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	// ...main code
}
```

### Run
pprof
```shell
make pprof source=<pprof endpoint>
# e.g
# make pprof source=http://localhost:6060/debug/pprof/heap
```

Heap memory profiling
```shell
make pprof-heap
```

Goroutine profiling
```shell
make pprof-goroutine
```

CPU profiling
```shell
make pprof-cpu
```

When these commands are executed, a browser will open, allowing you to view the pprof results. (Default URL: `localhost:9999`)

The pprof results are stored in the `$HOME/pprof` directory.
