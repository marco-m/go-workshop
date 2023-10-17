// Go Workshop © 2023 by Marco Molteni and contributors is licensed under CC BY-NC-SA 4.0
// https://creativecommons.org/licenses/by-nc-sa/4.0/

#import "@preview/polylux:0.3.1": *

#set page(
  paper: "presentation-16-9",
  fill: teal.lighten(90%),
  footer: [
      #set align(right)
      #set text(12pt)
      Go workshop - #logic.logical-slide.display()/#utils.last-slide-number
      //#utils.polylux-progress( ratio => [#calc.round(ratio * 100) #sym.percent])
    ]
)
#set text(size: 25pt, font: "Blogger Sans")

// #enable-handout-mode(true)

#show link: underline

////////////////////////

#polylux-slide[
  #set align(horizon + center)
  = Go workshop

  #v(2em)
  Marco Molteni

  2023-10-15

  #link("https://creativecommons.org/licenses/by-nc-sa/4.0/")[CC BY-NC-SA 4.0]
]

#polylux-slide[
== Getting started: before the workshop (1)
// sigh <before-workshop>

Please take the time to follow the instructions in the next slide. It will reduce frustration and will give you an editor that auto fixes problems as they appear (package imports, ...) and with code completion, code navigation, intellisense, ...

*Warning*: the majority of Go install guides on the Internet are wrong or old or misunderstand the `go.mod` way of doing things. Following the official links in the next slide is the fastest way to being productive.
]

#polylux-slide[
== Getting started: before the workshop (2)
1. Install the #link("https://go.dev/learn")[latest Go]\; do not use the one from your distribution.
2. Do not set any Go environment variable (`$GOPATH`, `$GOBIN`, ...).
3. Add two directories to your `$PATH`:
  - The path to the `bin` directory where you installed Go in step 1 (example: `/usr/local/go/bin`).
  - `$HOME/go/bin`. This is where `go install` will put executables.
]

#polylux-slide[
== Getting started: before the workshop (3)
4. Verify that `$PATH` is now set correctly:
  - The shell should find the `go` executable and the version should match what you installed in step 1:
    ```
    $ which go
    /usr/local/go/bin/go
    $ go version
    ```
  - Install the `gotestsum` test runner and verify that you can find it:
    ```
    $ go install gotest.tools/gotestsum@latest
    $ which gotestsum
    $HOME/go/bin/gotestsum
    ```
]

#polylux-slide[
== Getting started: before the workshop (4)
5. Take the time to #link("https://go.dev/doc/editors")[configure your editor] correctly.

Then, follow these two official tutorials:
6. #link("https://go.dev/doc/tutorial/getting-started")[Tutorial: Get started with Go]
7. #link("https://go.dev/doc/tutorial/create-module")[Tutorial: Create a Go module]
]

#polylux-slide[
== About this workshop

- Assumes no prior experience with Go.
- Assumes you followed the slides "Before the workshop".
- Assumes you already know how to program, maybe with a dynamically typed language.
- Is a work in progress, so incomplete.
- Is sometimes opinionated (that is: there could be other equally good ways of doing the same things).
]

#polylux-slide[
== About you
#pause
- What are your expectations from the workshop?
]

#polylux-slide[
== On learning something new
#pause
- Some people appreciate Go for its simplicity and accept its *defects*. Some don't. It is fine either way.
- Do not try to write in Go as if it were the language you are most familiar with.
- Instead, try to have the humility and patience to learn how to write idiomatic Go.
]

#polylux-slide[
== On what should be _easy_
#pause
Easy for me to write or easy for others to read?
#pause
- Readability (by your future self and your team members) is of the utmost importance.
- Being "easy to write" often clashes with readability.
- Go favours readability. Also if it can seem boring.

How?
#pause
- NO MAGIC.
- No monkey patching.
- No dynamic "changes" to the code.
]

#polylux-slide[
== Error handling
#pause
- Go has no exceptions, instead functions return multiple values, of which the last one, by convention, is the error.
- *Errors must be handled, immediately, _every_ time*.
- This can be tedious at first, but please just _swim with the flow_ and accept it.
- Your programs will be more robust and _sincere_: in the real world, errors happen all the time!
- You will also discover that testing code without exceptions is easier and explicit...
]

#polylux-slide[
== Error handling: example
#pause
#set text(size: 18pt)
```go
func enjoy() error {
    flavor, err := iceCream()
    if err != nil {
        return fmt.Errorf("enjoy: %s", err)
    }
    // Use flavor.
    // ...

    // All done, all OK.
    return nil
}

func iceCream() (string, error) {
    ...
}
```
]

#polylux-slide[
== Bubbling up errors
#pause
If we follow this approach, then the errors arrive up to the `main()` function:

#set text(size: 18pt)
```go
func main() {
  if err := run(); err != nil {
    fmt.Println("error:", err)
    os.Exit(1)
  }
}

func run() error {
   ...
}
```
This is also an answer to "how do I test the main function?" (more details later).
]

#polylux-slide[
  #set align(horizon + center)
  = Warmup: an hello word program
]

#polylux-slide[
== `helloworld`: a simplistic program

- Clone the workshop repo: #link("https://github.com/marco-m/go-workshop")[github.com/marco-m/go-workshop]
- `cd` to the #link("https://github.com/marco-m/go-workshop/tree/master/helloworld")[helloworld] directory
- The module structure is the simplest possible.
- Have a look at the README, run the commands explained there.
- Have a look at the code, kick the tires...
]

#polylux-slide[
  #set align(horizon + center)
  = `fruits`: a useful template to get started
]

#polylux-slide[
== `fruits`

- #link("https://github.com/marco-m/go-workshop/tree/master/fruits")[github.com/marco-m/go-workshop/fruits]
- Contains a *lot* of useful techniques and conventions for writing a command-line program.
- We will have a quick overview, but it is more for you to use as a reference in future projects.
]

#polylux-slide[
  #set align(horizon + center)
  = `loadmaster`: a useful command-line program
]

#polylux-slide[
== `loadmaster`: cumulative Concourse build times

Given a pipeline and a time window, calculate the total and average build time per job and display it, sorted per total time.

#set text(size: 20pt)
```
$ ./loadmaster build-time --pipeline=concourse
job                      count  average   total
dev-image                    7  2h45m17s  19h16m57s
resource-types-images       22  14m55s    5h27m59s
unit                         5  59m36s    4h58m2s
bosh-topgun-both             2  2h10m36s  4h21m11s
testflight                   4  49m27s    3h17m48s
bosh-topgun-runtime          2  1h12m36s  2h25m12s
bosh-topgun-core             2  59m18s    1h58m36s
...
```
]

#polylux-slide[
== Background: Concourse pipelines, jobs and builds
#pause
- A pipeline is a directed graph, where each node is a resource or a job.
- The graph can be connected or disconnected.
- A pipeline cannot be triggered!
  - Can only trigger a node (a resource check or a job build).
- Often a pipeline is a tree, and the root of the tree is a git resource.
  - Then triggering the git resource or the first job gives the impression of triggering the pipeline.
- A *build* is one execution of a *job*.
]

#polylux-slide[
== Backgroud: Concourse builds

#set text(size: 18pt)
```
$ fly -t developers builds --count=8

id   name             status     start          end      duration  team   created by
550  p1-mast/j13/123  started    2023-10-04@..  n/a      0s+       cloud  system
549  p2-mast/j1/5     started    2023-10-04@..  n/a      0s+       devs   system
548  p2-feat-5/j7/4   started    2023-10-04@..  n/a      0s+       cloud  system
530  p1-mast/j3/119   succeeded  2023-10-04@..  2023-..  3s        devs   system
529  p3-mast/j3/123   pending    n/a        ..  n/a      n/a       devs   bob@ex.org
525  p4-feat-9/j1/23  started    2023-10-04@..  n/a      1m1s+     cloud  system
519  p9-mast/j9/13    started    2023-10-04@..  n/a      1m31s+    cloud  system
518  p1-mast/j4/3     succeeded  2023-10-04@..  2023-..  25s       cloud  system
```
#set text(size: 20pt)
*Question*: What are:
- *id*:   #only(1)[?] #only(2)[`   global-build-Id`]
- *name*: #only(1)[?] #only(2)[`pipeline / job / relative-build-Id`]
]

#polylux-slide[
== What we are going to learn
#pause
- How to write a simple CLI program
- HTTP client
- JSON parsing
- Testing
- Testing HTTP clients with high fidelity
- ...
]

#polylux-slide[
== Back to `loadmaster`
#pause
Given a pipeline and a time window, calculate the total and average build time per job and display it, sorted per total time.

#set text(size: 20pt)
```
$ ./loadmaster build-time --pipeline=concourse
job                      count  average   total
dev-image                    7  2h45m17s  19h16m57s
resource-types-images       22  14m55s    5h27m59s
unit                         5  59m36s    4h58m2s
bosh-topgun-both             2  2h10m36s  4h21m11s
testflight                   4  49m27s    3h17m48s
bosh-topgun-runtime          2  1h12m36s  2h25m12s
bosh-topgun-core             2  59m18s    1h58m36s
...
```
]

#polylux-slide[
== `loadmaster`: first iteration
#set text(size: 18pt)
```
./bin/loadmaster -h
This program calculates statistics for Concourse
Usage: loadmaster [--server SERVER] [--team TEAM] [--timeout TIMEOUT] <command> [<args>]

Options:
  --server SERVER        Concourse server URL [default: https://ci.concourse-ci.org]
  --team TEAM            Concourse team [default: main]
  --timeout TIMEOUT      timeout for network operations (eg: 1h32m7s) [default: 5s]
  --version              display version and exit
  --help, -h             display this help and exit

Commands:
  build-time             calculate the cumulative build time taken by a pipeline

For more information visit FIXME https://example.org/...
```
]

#polylux-slide[
== `loadmaster`: getting started
- The exercise does not require authorization (no need to `fly login`).

- In the #link("https://github.com/marco-m/go-workshop")[github.com/marco-m/go-workshop] repo:
  - `loadmaster-skel`: the skeleton to use.
  - (`loadmaster`: solution to the exercise)
  - The tests are already there, you can use them as a guide.

- In the #link("https://github.com/concourse/concourse")[github.com/concourse/concourse] repo:
  - HTTP server routes: #link("https://github.com/concourse/concourse/blob/master/atc/routes.go")[concourse/atc/routes.go]
  - Looking at `fly` to understand which endpoint to use: #link("https://github.com/concourse/concourse/tree/master/fly/commands")[concourse/tree/master/fly/commands]
]

#polylux-slide[
== `loadmaster`: let's do it!

- Better to work in pairs.
- Start from loadmaster-skel, use the various links in the previous slides and try to write it.
- Feel free to ask if you have any doubt.
]

#polylux-slide[
== `loadmaster` going further 1
- Proposed additional flags:
#set text(size: 18pt)
```
    build-time --pipeline=NAME --day=DATE            (cumulative over all jobs)
    build-time --pipeline=NAME --day=DATE --per-job  (with details per job)
    build-time --pipeline=NAME --day=DATE --job=NAME (only that job)

    build-time --pipeline=NAME --from=DATETIME --to=DATETIME
```
\
#set text(size: 25pt)
- Understand response paging and how to navigate through them.
]

#polylux-slide[
== `loadmaster` going further 2
- Add flag to sort per average build time
- Add a column frequency, that shows (average) builds per day (requires to have solved pagination)
- Add a feature to report the top K most expensive jobs among ALL the pipelines of a Concourse team!, as usual, given a time window
- Add human time windows, for example `last-day`, `last-week`, `last-N-hours=N` ...
]

#polylux-slide[
== `loadmaster` going further 3
The time in human form is confusing to compare, because it can be:
```
123h32m7s
42s
```
while is should be something like:
```
123h 32m 07s
  0h 00m 42s
```
with minutes and seconds are always 2 digits, while the hours are the minimum digits for the value.
]

#polylux-slide[
== `loadmaster` going further 4
- Focusing on a given branch (eg `banana-master`) is not really representative of the Concourse resource consumption of a project.
- Instead, we should report cumulative time per project.
- This can be done by considering the pipeline *prefix*: `banana-master`, `banana-staging`, `banana-feat-1` and so on are all part of the *banana* project.
]

#polylux-slide[
#set align(horizon + center)
= Sources for learning Go and references
]

#polylux-slide[
== Online
More or less in suggested reading order:
- #link("https://go.dev/tour/welcome/1")[A tour of Go]
- #link("https://gobyexample.com/")[Go by Example]
- #link("https://go.dev/doc/effective_go")[Effective Go]
- #link("https://go.dev/doc/tutorial")[Go Tutorials]
- #link("https://go.dev/doc/")[Overview of the documentation]
- #link("https://go.dev/blog/")[The Go blog]

- #link("https://yourbasic.org/golang/")[Your Basic: Go] Very good!
- #link("https://bitfieldconsulting.com/golang/")[Bitfield Consulting: Go] very verbose but good, start from the oldest post and proceed
]

#polylux-slide[
== Books
- #link("https://lets-go.alexedwards.net/")[Let's Go] Build a web app in Go, step by step, server-side rendering with some JavaScript. Very good.
- #link("https://lets-go-further.alexedwards.net/")[Let's Go further] Build a web API in Go. Very good.
]

#polylux-slide[
== Videos
- All talks from Liz Rice, she is super. For example:
    - #link("https://www.youtube.com/watch?v=8fi7uSYlOdc")[Containers from scratch]
    - #link("https://www.youtube.com/watch?v=ZrpkrMKYvqQ")[Debuggers from scratch]
    - #link("https://www.youtube.com/watch?v=uBqRv8bDroc")[Beginner guide to eBPF programming with Go]
]

#polylux-slide[
== Free Trainings
- #link("https://exercism.org/tracks/go")[Exercism, the Go track] Good.
- #link("https://quii.gitbook.io/learn-go-with-tests/")[Learn Go with tests] TDD! very good.
- #link("https://gophercises.com/")[Gophercises] Did not take it but author is good, so I assume course is good.
]

#polylux-slide[
== Paying Trainings
- #link("https://www.ardanlabs.com/training/")[Ardan Labs] Very good. I followed the videos (the cheapest option).
- #link("https://codecrafters.io/")[CodeCrafters] Build your own Redis, Git, Docker, SQLite from scratch.
]
