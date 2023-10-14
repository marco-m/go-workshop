package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"

	"github.com/marco-m/go-workshop/fruits/internal"
	"github.com/marco-m/go-workshop/fruits/internal/parsley"
	"github.com/marco-m/go-workshop/fruits/pkg/banana"
)

// LinkerVersion must be set by the linker (see Taskfile).
var LinkerVersion = "unknown"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

type Args struct {
	ChangeMe int `help:"TODO"`
	//
	Version bool `help:"display version and exit"`
}

func (Args) Description() string {
	return "TODO This program does this and that."
}

func (Args) Epilogue() string {
	return "TODO For more information visit https://..."
}

func run(cmdLine []string) error {
	var args Args
	cli, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		return fmt.Errorf("init cli parsing: %s", err)
	}
	cli.MustParse(cmdLine)
	if args.Version {
		fmt.Println(internal.Version(LinkerVersion))
		return nil
	}

	fmt.Println("the color of parsley is", parsley.Color())

	lakatan := banana.Banana{}
	fmt.Println("the color of a banana is", lakatan.Color())
	if err := lakatan.Peel(); err != nil {
		return fmt.Errorf("run: peeling a banana: %s", err)
	}
	fmt.Println("the color of a peeled banana is", lakatan.Color())

	return nil
}
