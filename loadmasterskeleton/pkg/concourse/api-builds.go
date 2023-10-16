package concourse

import (
	"context"
	"fmt"
)

// BuildStatus taken from https://github.com/concourse/concourse/blob/master/atc/build.go
// Go doesn't have real enums; this is the simplest less bad we can do.
// Other approaches are possible, but I am not convinced of the trade-offs.
type BuildStatus string

const (
	StatusStarted BuildStatus = "started"
	// WRITEME
	//
	StatusUnknown BuildStatus = "unknown"
)

func ParseStatus(s string) (BuildStatus, error) {
	switch s {
	case "started":
		return StatusStarted, nil
	// WRITEME
	default:
		return StatusUnknown, fmt.Errorf("parse: unknown status: %q", s)
	}
}

func (status BuildStatus) String() string {
	return string(status)
}

type Build struct {
	Id int64 `json:"id"`
	// WRITEME
}

// UnmarshalJSON is normally not needed; function [json.Unmarshal] already knows
// what to do. In this case, we override the default behavior because we want to
// transparently parse fields "start_time" and "end_time", which on the wire are
// encoded as int64 (Unix time), to the Go-native [time.Time].
//
// For a detailed explanation, see
// - https://eli.thegreenplace.net/2019/go-json-cookbook/
// - https://choly.ca/post/go-json-marshalling/
func (bld *Build) UnmarshalJSON(data []byte) error {
	type Alias Build // Avoid infinite loop calling UnmarshalJSON.
	// WRITEME
	return nil
}

func (bld *Build) MarshalJSON() ([]byte, error) {
	type Alias Build // Avoid infinite loop calling MarshalJSON.
	// WRITEME
	return nil, fmt.Errorf("not implemented")
}

// ListPipelineBuilds returns the last builds for pipeline 'pipelineName'
// in team 'teamName'.
//
// From https://github.com/concourse/concourse/blob/master/atc/routes.go
// Path: "/api/v1/teams/:team_name/pipelines/:pipeline_name/builds"
// Method: "GET"
// Name: ListPipelineBuilds
func (cl *Client) ListPipelineBuilds(ctx context.Context,
	teamName, pipelineName string,
) ([]Build, error) {
	// WRITEME
	return nil, fmt.Errorf("not implemented")
}
