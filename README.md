# Go workshop

[![Build status](https://github.com/marco-m/go-workshop/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/marco-m/go-workshop/actions)

This repo contains the materials for a WIP Go workshop.

Directory layout:

- [slides](./slides) The slides of the workshop.

- [helloworld](./helloworld) The classic helloworld program.
- [fruits](./fruits) Example of a small CLI program and test patterns.
- [loadmaster](./loadmaster) Calculate the cumulative Concourse build times.
- [loadmasterskeleton](./loadmasterskeleton) Just the skeleton for loadmaster, to write yourself!

## gonew templates

Optionally, each directory is a project template, with its own Go module, and can be obtained with the [gonew](https://go.dev/blog/gonew) tool.

If for example you want to create a Go module using the module `fruits` as starting point:

- Install `gonew`: `go install golang.org/x/tools/cmd/gonew@latest`
- Run `gonew`:
```
$ gonew github.com/marco-m/go-workshop/fruits github.com/$MYORG/$MYPROJECT
$ cd $MYPROJECT
```
- Init the git repo.
- Add `bin/` to `.gitignore` (see Taskfile for why).
- Run `task --list` to see all the targets.
- Run `task install:deps`.
- Inspect and adapt the `LICENSE` file.
- Inspect and adapt the `Taskfile.yml` file.
- Inspect and adapt all the other files.

## Requirements

- [Go](https://go.dev/) version >= 1.21
- [Taskfile](https://taskfile.dev)

## Licenses

- The source code is licensed under the [MIT LICENSE](SOURCE.LICENSE).
- The slides are licensed under Creative Commons [Attribution-NonCommercial-ShareAlike 4.0 International (CC BY-NC-SA 4.0)](https://creativecommons.org/licenses/by-nc-sa/4.0/).
