package concourse

import (
	"context"
	"fmt"
	"net/http"
)

// Client is a minimal client for the Concourse HTTP API.
// Do not instantiate directly; instead use NewClient.
type Client struct {
	Server     string       // Mandatory.
	HttpClient *http.Client // Optional; to be overridden in tests.
}

// NewClient instantiates a Concourse [Client].
func NewClient(opts Client) (*Client, error) {
	// WRITEME
	return nil, fmt.Errorf("not implemented")
}

func get(ctx context.Context, hclient *http.Client, urlo string) ([]byte, error) {
	// WRITEME
	return nil, fmt.Errorf("not implemented")
}
