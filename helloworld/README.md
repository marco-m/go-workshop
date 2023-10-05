# helloworld

This is an absolute minimum example of Go.

It doesn't use dependencies and is built and tested directly with the `go` tool.

Contrary to the other examples in this repo, it uses an almost flat directory layout.

It is meant to be observed commit per commit on the `helloworld` branch.

```
.
├── README.md
├── go.mod
├── go.sum
├── main.go
└── main_test.go
```

## Build

It is not mandatory to put the build artifacts in the `./bin` directory; we do this because we have `bin/` in `.gitignore` and this is handy.

```
$ mkdir bin
$ go build -C bin ..
```

## Test

```
$ go test -v .
```
