package main

import (
	"fmt"
	"os"

	"github.com/marco-m/go-project-templates/helloworld/hello"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func run(cmdLine []string) error {
	// This is silly; used only to show some more testing.
	if len(cmdLine) > 0 {
		return fmt.Errorf("unexpected: %q", cmdLine)
	}
	fmt.Println(hello.Greet("bob"))
	return nil
}
