package main

import (
	"fmt"
	"os"

	"github.com/marco-m/go-project-templates/helloworld/hello"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println(hello.Greet("bob"))
	return nil
}
