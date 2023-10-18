package main

import (
	"fmt"
	"os"

	"github.com/marco-m/go-workshop/http-app/pkg/app"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Println("app: error:", err)
		return 1
	}
	return 0
}
