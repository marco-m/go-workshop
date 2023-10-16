package concourse

import (
	"context"
	"errors"
	"fmt"
	"io"
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
	client := opts
	if client.Server == "" {
		return nil, errors.New("concourse.NewClient: empty field Server")
	}
	if client.HttpClient == nil {
		client.HttpClient = &http.Client{}
	}

	return &client, nil
}

func get(ctx context.Context, hclient *http.Client, urlo string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlo, nil)
	if err != nil {
		return nil, fmt.Errorf("get: new request: %s", err)
	}
	resp, err := hclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get: do: %s", err)
	}
	defer resp.Body.Close()

	body, errBody := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if errBody != nil {
			return nil, fmt.Errorf("get: status code: %d (%s)", resp.StatusCode, errBody)
		}
		return nil, fmt.Errorf("get: status code: %d (%s)", resp.StatusCode, string(body))
	}
	if errBody != nil {
		return body, fmt.Errorf("get: read body: %s", errBody)
	}

	return body, nil
}
