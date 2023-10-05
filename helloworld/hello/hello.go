package hello

import "fmt"

func Greet(who string) string {
	return fmt.Sprintf("hello %s", who)
}
