package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/alexflint/go-arg"

	"github.com/marco-m/go-workshop/http-app/internal"
)

// LinkerVersion must be set by the linker (see Taskfile).
var LinkerVersion = "unknown"

type args struct {
	Listen string `help:"the TCP listen address, in the form host:port. An empty host means all interfaces"`
}

func (args) Version() string {
	return internal.Version(LinkerVersion)
}

func (args) Description() string {
	return "WRITEME"
}

func (args) Epilogue() string {
	return "For more information visit FIXME https://example.org/..."
}

type application struct {
	cfg args
	//
	version string
	log     *slog.Logger
}

func Run(cmdLine []string) error {
	app := application{
		cfg: args{
			Listen: ":80",
		},
		version: internal.Version(LinkerVersion),
		log:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	argParser, err := arg.NewParser(arg.Config{}, &app.cfg)
	if err != nil {
		return fmt.Errorf("init cli parsing: %s", err)
	}
	argParser.MustParse(cmdLine)

	app.log.Info("", "state", "starting", "listen", app.cfg.Listen)
	return http.ListenAndServe(app.cfg.Listen, app.routes())
}
