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

## Hello, world

*Declaring variables, constants, if/else statements, switch, write your first
go program and write your first test. Sub-test syntax and closures.*

### First App

:shipit: Go apps need to live in [`$GOPATH`](https://golang.org/doc/gopath_code.html)?
```bash
mkdir -p $GOPATH/src/github.com/$USER/hello
```

:fire: figured out how to build apps in a local dir instead of `$GOPATH` with modules

### First Test

[:ship: 23ebc5c](https://github.com/arafatm/learn-go-with-tests/commit/23ebc5c)
Our first test
- `t.Errorf` fails test and prints a message

### Go Docs

Run Go docs
```bash
godoc -http :8000
```

- http://localhost:8000/pkg  to see installed packages
- http://localhost:8000/pkg/testing/ to see testing package

### Args & constants

[:ship: 58a84c9](https://github.com/arafatm/learn-go-with-tests/commit/58a84c9)
method arguments

[:ship: 86a5030](https://github.com/arafatm/learn-go-with-tests/commit/86a5030)
Using constants `const englishHelloPrefix = Hello, `

[:ship: be312af](https://github.com/arafatm/learn-go-with-tests/commit/be312af)
Go does not have default values for method args

### Test Helpers

[:ship: 79c79e9](https://github.com/arafatm/learn-go-with-tests/commit/79c79e9)
refactor tests with a `t.Helper()`

### Switch statement

[:ship: 803e3f4](https://github.com/arafatm/learn-go-with-tests/commit/803e3f4)
Add language support with `if` statements

[:ship: 5e3e90f](https://github.com/arafatm/learn-go-with-tests/commit/5e3e90f)
refactor `if` statements with `switch`

## Examples

*Further Explore function declaration syntax and learn new ways to improve the
documentation of your code.*

:neckbeard: I'm using the suggestions for setting up Go projects here
https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project.
Note the `package` and `module` declarations.

[:ship: 184e090](https://github.com/arafatm/learn-go-with-tests/commit/184e090)
First `adder` and experiment with Go package and module naming

[:ship: 0d3913e](https://github.com/arafatm/learn-go-with-tests/commit/0d3913e)
sample `Example` that can be verified by running `godoc`
- [Examples](https://blog.golang.org/examples) are testable code blocks that
  can be displayed as **package documentation**

You can also `go test -v` to run the tests and examples 

[:ship: 0b5ef61](https://github.com/arafatm/learn-go-with-tests/commit/0b5ef61)
go-linter wants functions to have comment describing them

## Iteration

*Learn about for and benchmarking.*

In Go there are **no while, do, until** keywords, you can only use `for`

xxx

## Arrays and slices

*Learn about arrays, slices, len, varargs, range and test coverage.*

## Structs, methods & interfaces

*Learn about struct, methods, interface and table driven tests.*

## Pointers & errors

*Learn about pointers and errors.*

## Maps

*Learn about storing values in the map data structure.*

## Dependency Injection

*Learn about dependency injection, how it relates to using interfaces and a
primer on io.*

## Mocking

*Take some existing untested code and use DI with mocking to test it.*

## Concurrency

*Learn how to write concurrent code to make your software faster.*

## Select

*Learn how to synchronise asynchronous processes elegantly.*

## Reflection

*Learn about reflection*

## Sync

*Learn some functionality from the sync package including WaitGroup and Mutex*

## Context

*Use the context package to manage and cancel long-running processes*

## Intro to property based tests

*Practice some TDD with the Roman Numerals kata and get a brief intro to
property based tests*

## Maths

*Use the math package to draw an SVG clock*
