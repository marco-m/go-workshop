// This file runs tests using the 'testscript' package.
// To understand, see:
// - https://github.com/rogpeppe/go-internal
// - https://bitfieldconsulting.com/golang/test-scripts

package loadmaster_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/marco-m/go-workshop/loadmaster/pkg/loadmaster"
)

func TestMain(m *testing.M) {
	// The commands map holds the set of command names, each with an associated
	// run function which should return the code to pass to os.Exit.
	// When [testscript.Run] is called, these commands are installed as regular
	// commands in the shell path, so can be invoked with "exec".
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"loadmaster":     loadmaster.Main,
		"build-time-vcr": makeBuildTimeVcr(),
	}))
}

func makeBuildTimeVcr() func() int {
	return func() int {
		ctx := context.Background()
		rec, err := recorder.NewWithOptions(&recorder.Options{
			CassetteName:       "list-pipeline-builds-short",
			Mode:               recorder.ModeReplayOnly,
			RealTransport:      http.DefaultTransport,
			SkipRequestLatency: true,
		})
		if err != nil {
			fmt.Println(err)
			return 1
		}
		global := loadmaster.Global{
			Server:     "https://ci.concourse-ci.org",
			Team:       "main",
			HttpClient: rec.GetDefaultClient(),
		}
		local := &loadmaster.BuildTimeCmd{Pipeline: "concourse"}
		if err := loadmaster.CmdBuildTime(ctx, global, local); err != nil {
			fmt.Println(err)
			return 1
		}
		return 0
	}
}

func TestScriptLoadmaster(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
