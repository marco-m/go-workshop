package concourse

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// BuildStatus taken from https://github.com/concourse/concourse/blob/master/atc/build.go
// Go doesn't have real enums; this is the simplest less bad we can do.
// Other approaches are possible, but I am not convinced of the trade-offs.
type BuildStatus string

const (
	StatusStarted   BuildStatus = "started"
	StatusPending   BuildStatus = "pending"
	StatusSucceeded BuildStatus = "succeeded"
	StatusFailed    BuildStatus = "failed"
	StatusErrored   BuildStatus = "errored"
	StatusAborted   BuildStatus = "aborted"
	//
	StatusUnknown BuildStatus = "unknown"
)

func ParseStatus(s string) (BuildStatus, error) {
	switch s {
	case "started":
		return StatusStarted, nil
	case "pending":
		return StatusPending, nil
	case "succeeded":
		return StatusSucceeded, nil
	case "failed":
		return StatusFailed, nil
	case "errored":
		return StatusErrored, nil
	case "aborted":
		return StatusAborted, nil
	default:
		return StatusUnknown, fmt.Errorf("parse: unknown status: %q", s)
	}
}

func (status BuildStatus) String() string {
	return string(status)
}

type Build struct {
	Id           int64       `json:"id"`
	TeamName     string      `json:"team_name"`
	Name         string      `json:"name"`
	Status       BuildStatus `json:"status"`
	ApiUrl       string      `json:"api_url,omitempty"`
	JobName      string      `json:"job_name"`
	PipelineId   int64       `json:"pipeline_id"`
	PipelineName string      `json:"pipeline_name"`
	StartTime    time.Time   `json:"start_time"` // on the wire: int64
	EndTime      time.Time   `json:"end_time"`   // on the wire: int64
	CreatedBy    string      `json:"created_by,omitempty"`
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
	aux := &struct {
		StartTime int64 `json:"start_time"`
		EndTime   int64 `json:"end_time"`
		*Alias
	}{
		Alias: (*Alias)(bld),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bld.StartTime = time.Unix(aux.StartTime, 0).UTC()
	bld.EndTime = time.Unix(aux.EndTime, 0).UTC()

	return nil
}

func (bld *Build) MarshalJSON() ([]byte, error) {
	type Alias Build // Avoid infinite loop calling MarshalJSON.
	return json.Marshal(&struct {
		StartTime int64 `json:"start_time"`
		EndTime   int64 `json:"end_time"`
		*Alias
	}{
		StartTime: bld.StartTime.Unix(),
		EndTime:   bld.EndTime.Unix(),
		Alias:     (*Alias)(bld),
	})
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
	urlo, err := url.JoinPath(cl.Server,
		"/api/v1/teams/", teamName, "/pipelines", pipelineName, "/builds")
	if err != nil {
		return nil, err
	}
	body, err := get(ctx, cl.HttpClient, urlo)
	if err != nil {
		return nil, fmt.Errorf("ListPipelineBuilds: url: %s: %s", urlo, err)
	}
	var builds []Build
	if err := json.Unmarshal(body, &builds); err != nil {
		return nil, err
	}
	return builds, nil
}
