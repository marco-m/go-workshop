// This file runs tests using the 'testscript' package.
// To understand, see:
// - https://github.com/rogpeppe/go-internal
// - https://bitfieldconsulting.com/golang/test-scripts

package loadmaster_test

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"

	"github.com/marco-m/go-workshop/loadmaster/pkg/loadmaster"
)

func TestMain(m *testing.M) {
	// The commands map holds the set of command names, each with an associated
	// run function which should return the code to pass to os.Exit.
	// When testscript.Run is called, these commands are installed as regular
	// commands in the shell path, so can be invoked with "exec".
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"loadmaster": loadmaster.Main,
	}))
}

func TestScriptLoadmaster(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
