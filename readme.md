# Learn Go With Tests

https://quii.gitbook.io/learn-go-with-tests/

## Install Go

*Set up environment for productivity.*

:shipit: Install Go && setup env 
```bash
brew install go

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

mkdir -p $GOPATH/src $GOPATH/pkg $GOPATH/bin

go version
```

:shipit: Go 1.11 introduced
[Modules](https://github.com/golang/go/wiki/Modules) as alternative to `GOPATH`
```bash
go mod init <modulepath>
```

:shipit: In `go.mod` 
```
module cmd

go 1.14

require (
        github.com/google/pprof v0.0.0-20190515194954-54271f7e092f
        golang.org/x/arch v0.0.0-20190815191158-8a70ba74b3a1
        golang.org/x/tools v0.0.0-20190611154301-25a4f137592f
)
```


:shipit: Help for `go mod`
```bash
go help mod
go help mod init
```


:shipit: If using Visual Studio Code
```bash
code --install-extension golang.go
```

:shipit: Install Delve for debugging
```bash
go get -u github.com/go-delve/delve/cmd/dlv
```

:shipit: Go Linter
```bash
brew install golangci/tap/golangci-lint
```

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world using Go mods 


[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

$msg

xxx

## Hello, world

*Declaring variables, constants, if/else statements, switch, write your first
go program and write your first test. Sub-test syntax and closures.*

:shipit: Go apps need to live in [`$GOPATH`](https://golang.org/doc/gopath_code.html)?
```bash
mkdir -p $GOPATH/src/github.com/$USER/hello
```

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Integers

*Further Explore function declaration syntax and learn new ways to improve the
documentation of your code.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Iteration

*Learn about for and benchmarking.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Arrays and slices

*Learn about arrays, slices, len, varargs, range and test coverage.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Structs, methods & interfaces

*Learn about struct, methods, interface and table driven tests.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Pointers & errors

*Learn about pointers and errors.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Maps

*Learn about storing values in the map data structure.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Dependency Injection

*Learn about dependency injection, how it relates to using interfaces and a
primer on io.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Mocking

*Take some existing untested code and use DI with mocking to test it.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Concurrency

*Learn how to write concurrent code to make your software faster.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Select

*Learn how to synchronise asynchronous processes elegantly.*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Reflection

*Learn about reflection*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Sync

*Learn some functionality from the sync package including WaitGroup and Mutex*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Context

*Use the context package to manage and cancel long-running processes*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Intro to property based tests

*Practice some TDD with the Roman Numerals kata and get a brief intro to
property based tests*

[:ship: 700dc62](https://github.com/arafatm/learn-go-with-tests/commit/700dc62)
hello world
```bash
go run hello.go
```

[:ship: 1c3e6d4](https://github.com/arafatm/learn-go-with-tests/commit/1c3e6d4)
hello world

[:ship: 27e2a58](https://github.com/arafatm/learn-go-with-tests/commit/27e2a58)
hello world

[:ship: 27bf756](https://github.com/arafatm/learn-go-with-tests/commit/27bf756)
testing

[:ship: e1afda8](https://github.com/arafatm/learn-go-with-tests/commit/e1afda8)
testing 2

xxx

## Maths

*Use the math package to draw an SVG clock*

xxx
