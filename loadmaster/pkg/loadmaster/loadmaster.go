package loadmaster

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/alexflint/go-arg"

	"github.com/marco-m/go-workshop/loadmaster/internal"
)

// LinkerVersion must be set by the linker (see Taskfile).
var LinkerVersion = "unknown"

type Args struct {
	Global
	//
	BuildTime *BuildTimeCmd `arg:"subcommand:build-time" help:"calculate the cumulative build time taken by a pipeline"`
}

type Global struct {
	Server  string        `help:"Concourse server URL"`
	Team    string        `help:"Concourse team"`
	Timeout time.Duration `help:"timeout for network operations (eg: 1h32m7s)"`
	//
	Version bool `help:"display version and exit"`
}

func (Args) Description() string {
	return "This program calculates statistics for Concourse"
}

func (Args) Epilogue() string {
	return "For more information visit FIXME https://example.org/..."
}

type BuildTimeCmd struct {
	Pipeline string `arg:"required"`
}

func Main() int {
	if err := run(os.Args[1:]); err != nil {
		fmt.Println("error:", err)
		return 1
	}
	return 0
}

func run(cmdLine []string) error {
	args := Args{
		Global: Global{
			Server:  "https://ci.concourse-ci.org",
			Team:    "main",
			Timeout: 5 * time.Second,
		},
	}
	argParser, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		return fmt.Errorf("init cli parsing: %s", err)
	}
	argParser.MustParse(cmdLine)
	if args.Version {
		fmt.Println(internal.Version(LinkerVersion))
		return nil
	}
	if argParser.Subcommand() == nil {
		argParser.Fail("missing subcommand")
	}

	ctx, cancel := context.WithTimeout(context.Background(), args.Timeout)
	defer cancel()

	switch {
	case args.BuildTime != nil:
		return cmdBuildTime(ctx, args.Global, args.BuildTime)
	default:
		return fmt.Errorf("internal error: unwired subcommand: %s", argParser.Subcommand())
	}
}
