=== Idea 3: Google Chat bot
Adaptation of my multigateway project.
Ideally, we would like to be able to offer the bot service by connecting to the Google Chat API as one can do for Telegram (and probably Zulip). Unfortunately Gchat forces to listen to be connected to, why why why? On the other hand, if we can use a service like ngrok, then this would be an interesting exercise for the workshop.
I had a look once more. Google documentation, and the Google chat architecture, make me vomit.

Go SDK
- `go get google.golang.org/api/chat/v1`
- code: https://github.com/googleapis/google-api-go-client/tree/main/chat/v1

=== Idea 4: Zulip chat bot
Adaptation of my multigateway project.
This is maybe reasonable?
Can either ask before to create an account on the Cloud team instance (ah it seems it disappeared???) or create another one for the workshop, or just share Poldoz?
Mah, at the end, we risk spending time trying to understand the Zulip API?
In comparison, the Concourse API seems way simpler, especially because we only want to query?
Maybe I can avoid spending time understanding the Zulip API if I suggest to use the package https://github.com/ifo/gozulipbot ?

Two ways to interact with bots:
- in a public (or private) stream
- in a DM (Direct Message)
for what we want to do, we want to support only interactions in a stream

- Add bot to organization
- DM bot and invite to a given stream
- in that stream, send a message to the bot with the classic `@` sign, for example
```
@ciccio weather
```
where the `@ciccio` can appear anywhere.

=== Idea 5: permuted index (KWIC)
Pros: no network needed
Cons: more boring than fiddling with an API over the network?

- "permuted index"? (from sqlite web site)

> A permuted index is one in which each phrase is indexed by every word in the phrase.

Also called [Key Word in Context (wikipedia)](https://en.wikipedia.org/wiki/Key_Word_in_Context)
> A KWIC index is formed by sorting and aligning the words within an article title to allow each word (except the stop words) in titles to be searchable alphabetically in the index. It was a useful indexing method for technical manuals before computerized full text search became common.

A KWIC index is a special case of a _permuted index_

David L. Parnas uses a KWIC Index as an example on how to perform modular design in his paper _On the Criteria To Be Used in Decomposing Systems into Modules_.





== Suggested skeleton for a Go project
- There are various conventions for the directory layout; this is the one that I find clearer.
- There are various approaches for command-line parsing, logging, testing. This is the one that I found the best.
- There are various approaches to build a Go project. I will show one.




== Project directory layout
[Introduction to `gonew`](https://go.dev/blog/gonew)
Actually `gonew` in its simplicity is not that bad. Mah, really? How do I change names for the executables? It seems I cannot?


```
README.md
cmd/
  tool1/
    main.go
  tool2/
    main.go
internal/
  parsley/
    foo.go
pkg/
  banana/
  mango/
```

main and run
Go-arg for parsing
First show go run and go build
Then skeleton Taskfile? Or my cook approach actually?
Then http package
version with linker flags since go support is still abandoned?
gotestsum ?
code coverage
browser task target
goreleaser
golangci-lint
unit tests
integration tests, skip if not env var
auto-update
internal/testscript!!!


Then I would like to show also testing... top would be the two levels, acceptance and unit, but risks becoming confusing too much meat on the fire and artificial? Works only because prepared before???


=== Testing: getting things straight
- Monkey patching is impossible in Go and is IMO a bad idea no matter the language.
- Mocking (or in general using test doubles) is naturally possible in Go, with classic DI (dependency injection).



=== When possible, do not mock
instead:
- Use [net/http/httptest](https://pkg.go.dev/net/http/httptest).
- Use testdata fixtures.
- Use virtual FS [testing/fstest/MapFS](https://pkg.go.dev/testing/fstest).
- Use testcontainers.
- Use a powerful DSL to test the I/O of an executable with [rogpeppe/go-internal](https://github.com/rogpeppe/go-internal).


=== Why avoiding mocking?
+ What is the most important question to keep asking oneself when writing a test?
+ → *WHAT am I testing???*
+ With a mock, it is very easy to loose track of what one is testing: the SUT or... the mock?
+ Is the mock representative of what it is mocking?
+ Is the SUT interacting according to the behavior of the _real_ collaborator or according to the... mock???


=== I am not convinced
- [Martin Fowler, Mocks aren't stubs](https://martinfowler.com/articles/mocksArentStubs.html)
- Are you a _classicist_ or a _mockist_?


=== Example of too much mocking
- The example in pipeline-setter that fills a public error type from go-github is a very good example!
- The one that I just found is even better, the list comments, create comment sequence...


=== Mocking with Dependency Injection
- Sometimes one _needs_ to mock.
- If you need to mock, use _classic_ DI, no _magic_ DI frameworks.
- DI can be implemented with `interfaces` or with first-class functions.


- video advanced testing mitchel hashimoto
- helper: the less worse is [go-quicktest/qt](https://github.com/go-quicktest/qt)
- Table-driven tests, split success and failure cases, try to use a setup function, do not cede to the trpbof stuffing everything into a test struct!!! Example of spkitting in two, so that one test is simpler and doesn't need the network at all (and so no need of testhttp) real life example from fly helper
- How to use t.Parallel()
- go.dev/blog/subtests

Two sources for mocking and avoiding monkey patching
- https://github.com/gotestyourself/gotest.tools/wiki/Test-Doubles-And-Patching
- leanr go with tests, two sections:
    - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
    - https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/working-without-mocks





== Acceptance tests
Links from [Learn Go with tests/Scaling acceptance tests](https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/scaling-acceptance-tests):
- [David Farley, How to Write Acceptance Tests](https://www.youtube.com/watch?v=JDD5EEJgpHU). Short and to the point.
- [Nat Pryce, End to end functional tests that can run in milliseconds](https://www.youtube.com/watch?v=Fk4rCn4YLLU). More verbose and confusing, but good insights.



=== How to test sleeping without sleeping
An example of dependency injection, function closures and hand-written spies.



=== Testing: high-fidelity HTTP reply
There are 2 libraries:
- https://github.com/seborama/govcr
- https://github.com/dnaeon/go-vcr Seems way simpler, for sure less code, so to prefer?
    - Exactly what I want: [using go-vcr with google/go-github and oauth2 (#59)](https://github.com/dnaeon/go-vcr/issues/59)
    - Blog from the author: http://dnaeon.github.io/testing-http-interactions-in-go/
Which is best?

I remember that there was also another library, with a different name (not containing VCR in its name)?
Ah yes, https://goreplay.org/, https://github.com/buger/goreplay. Seems powerful but also complex, I would say not needed for what I want to do.




What about https://quii.gitbook.io/learn-go-with-tests ?



Later we will see more details, such as testing loggers, command-line parsing, cancellation events (`context.Context`) and so on.




== project layout
show options, explain the name taken from directory
Command-line parsing with go-arg
== Go runtime: GG and pointers
== go.mod and go.work, updating deps
[Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)

== 3 gotchas
- pointers make a difference (make example with struct method and modifying struct fields)
- loop capture with closures (https://go.dev/blog/loopvar-preview)
- errors must always be handled: no exceptions: program keeps going until it burns in flames

== OO?
- No inheritance.
- You can survive and live happy without OOP.
- structs with methods.
-

== do I need multiple go installed?

== what to avoid
- global variables
- Panic
- Log.fatal log.panic
- Instead of log, use log/slog
- os.Exit (except, obviously, from the `main()` function)

== documenting Go programs


== old stuff on the Internet you must not follow
GOPATH
GO11 or similar

== formatting
gofmt
go fmt
gofumt

== interfaces

== some beautiful stdlib interfaces
... the 1 method ones!


== errors and error handling
https://bitfieldconsulting.com/golang/testing-errors
https://bitfieldconsulting.com/golang/comparing-errors
https://bitfieldconsulting.com/golang/wrapping-errors
https://earthly.dev/blog/golang-errors/

errors.As subtleties

```
      return fmt.Errorf("workFn: %w",
        BananaResponseError{
          Response: &BananaResponse{
            Amount:  42,
            variety: "space yellow",
          }})
```
needs
```
    var errBananaResponse BananaResponseError
    if errors.As(err, &errBananaResponse) {
      return retry.SoftFail
    }
```
while
```
      return fmt.Errorf("workFn: %w",
        &BananaResponseError{
          Response: &BananaResponse{
            Amount:  42,
            variety: "space yellow",
          }})
    }
```
needs
```
    var errBananaResponse *BananaResponseError
    if errors.As(err, &errBananaResponse) {
      return retry.SoftFail
    }
```
== foo

 the equivalent of "not visible from outside" is `internal`, see [https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface).

In this, Go is a bit messy:

- any package must be in its directory
- `pkg` is purely a convention, many repos do not use it, they have `project/foo` instead of `project/pkg/foo`
- I think that `pkg` helps, because say in a web project, directory `/project/assets` could be web assets or a Go package with name `assets`...
- on the other hand, `internal` is enforced by the compiler: project/internal/foo CANNOT be imported by another module (module == to simplify == a Go project)
- other directory name based purely on convention is `cmd`, used especially when a project has more than one executable: `project/cmd/banana`, `project/cmd/mango`
- there is also an impact on how easy it is to use `go install`: if a project has the `main` package at its root, then `go install github.com/foo/bar` will magically make available the executable `bar` under `HOME/go/bin`.
- if on the other hand bar is `foo/cmd/bar`, then `go install github.com/foo/bar` will give a error to stare at without help to understand; after a while a light bulb goes on and you do... `go install github.com/foo/cmd/bar`

https://ieftimov.com/posts/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/
== http
do not need any http library. Just use the stdlib, both for clients and servers.

== http server, http app

== command-line parsing

== logging

== advanced

=== generics

=== goroutines and channels

=== context



// SIGH #set page(fill: orange)
#polylux-slide[
  #set align(horizon + center)
  = Backup slides
]

#polylux-slide[
== Idea 1: Concourse find all pipelines referencing a given repository
Proposed command:
```
aircargo cross-reference --list-pipelines --referenced REPO
<list of pipelines referencing that REPO>
(other variants can be considered...)
```

- Concourse has multiples teams. Each pipeline uses one or more git repositories. Each git repository can be used by multiple pipelines (that is: many to many relationship).
- Given a team, query Concourse to make a list of the pipelines associated to each git repository.
- Example:
    - repo-1:
        - pl-A
    - repo-2:
        - pl-A
        - pl-B
    - ...
]
