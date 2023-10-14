# fruits CLI tool

This is the README for the fruits CLI tool sample project.

It shows:
- A possible approach to Go module layout.
- Command-line parsing with [go-arg](https://github.com/alexflint/go-arg).
- Design errors in the API and their impact on testing (see `banana.go`).
- Table-driven tests and their relation with xUnit setup and teardown.
- Testing with [go-quicktest](https://github.com/go-quicktest/qt).
- Focus on testing public API (e.g. banana_test.go); test private API (e.g. banana_private_test.go) only when needed.
- Using a test spy to simulate time.Sleep().
- How to skip integration tests.
- Basic [Fuzz testing](https://go.dev/security/fuzz/).
- The [gotestsum](https://github.com/gotestyourself/gotestsum) test runner.
- The [internal](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface) directory.
- How to embed version information in the binary (see `Taskfile.yml` and `internal/version.go`)
- How to use [Taskfile](https://taskfile.dev) for automation.

```
.
├── README.md
├── Taskfile.yml
├── bin/             <= Created by Taskfile; build artifacts will be here.
├── cmd/
│   └── fruits/
│       └── main.go  <= Contains the main() function.
├── go.mod
├── go.sum
├── internal/
│   ├── parsley/     <= Example internal package.
│   │   └── parsley.go
│   └── version.go   <= Logic to report build version; see also the Taskfile.
└── pkg
    └── banana       <= Example package with tests.
        ├── banana.go
        ├── banana_fuzz_test.go
        ├── banana_private_test.go
        └── banana_test.go
```

## Taskfile

Run `task --list` to see the available targets.

## Install tool dependencies

```
$ task install:deps
```

## Build

```
$ task build
or
$ task build && ./bin/fruits --help
```

## Test

```
$ task --list | grep test:
* test:unit:  Run the unit tests. Some tests will be listed as "Skipped".
* test:all:   Run all the tests. No tests will be skipped.

* browser:    Show code coverage in browser (usage: task test:<subtarget> browser)
* test:fuzz:  Run all the fuzz tests. Interrupt with Ctrl-C.
```

### Unit tests

This shows the skipping of tests needing special setup:

```
$ task test:unit
task: [test:unit] gotestsum -- -count=1 -coverprofile=bin/coverage.out ./...
∅  internal/parsley
∅  internal
∅  cmd/fruits (1ms)
✓  pkg/banana (177ms) (coverage: 65.4% of statements)

=== Skipped
=== SKIP: pkg/banana TestBananaIntegration (0.00s)
    banana_test.go:107: Skipping integration test: missing env var BANANA_TEST_TOKEN

DONE 11 tests, 1 skipped in 0.830s
```

### All tests

Compare the output with the unit tests:

```
task test:all
task: [test:all] gotestsum -- -count=1 -coverprofile=bin/coverage.out ./...
∅  internal/parsley
∅  internal
∅  cmd/fruits (1ms)
✓  pkg/banana (153ms) (coverage: 65.4% of statements)

DONE 11 tests in 0.536s
```

In this case, the coverage is the same as the unit tests because test `TestBananaIntegration` is empty.
