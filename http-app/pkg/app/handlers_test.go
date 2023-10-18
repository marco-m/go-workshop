// Based on example from the book "Let's Go" by Alex Edwards,
// chapter "End-to-end testing"

package app

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-quicktest/qt"
)

func TestPingHandler(t *testing.T) {
	app := &application{
		log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/ping")
	qt.Assert(t, qt.IsNil(err))
	defer resp.Body.Close()
	qt.Assert(t, qt.Equals(resp.StatusCode, http.StatusOK))

	body, err := io.ReadAll(resp.Body)
	qt.Assert(t, qt.IsNil(err))

	body = bytes.TrimSpace(body)
	qt.Assert(t, qt.Equals(string(body), "OK"))
}
