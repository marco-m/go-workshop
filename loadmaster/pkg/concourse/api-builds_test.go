package concourse_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/go-quicktest/qt"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/marco-m/go-workshop/loadmaster/pkg/concourse"
)

func TestClient_ListPipelineBuilds(t *testing.T) {
	// Arrange recorder.
	rec, err := recorder.NewWithOptions(&recorder.Options{
		CassetteName:       "testdata/list-pipeline-builds-short",
		Mode:               recorder.ModeRecordOnce,
		RealTransport:      http.DefaultTransport,
		SkipRequestLatency: true,
	})
	qt.Assert(t, qt.IsNil(err))

	// Arrange SUT.
	const teamName = "main"
	const pipelineName = "concourse"
	concourseClient, err := concourse.NewClient(concourse.Client{
		Server:     "https://ci.concourse-ci.org",
		HttpClient: rec.GetDefaultClient(),
	})
	qt.Assert(t, qt.IsNil(err))
	ctx := context.Background()

	// Act.
	got, err := concourseClient.ListPipelineBuilds(ctx, teamName, pipelineName)

	// Teardown
	qt.Assert(t, qt.IsNil(rec.Stop())) // Needed to actually save the cassette on first run.

	// Assert.
	qt.Assert(t, qt.IsNil(err))
	want := []concourse.Build{
		{
			Id:           214953806,
			TeamName:     "main",
			Name:         "575",
			Status:       "started",
			ApiUrl:       "/api/v1/builds/214953806",
			JobName:      "dev-image",
			PipelineId:   24,
			PipelineName: "concourse",
			StartTime:    time.Date(2023, time.October, 16, 14, 7, 45, 0, time.UTC),
			EndTime:      time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			CreatedBy:    "",
		},
		{
			Id:           214953805,
			TeamName:     "main",
			Name:         "402",
			Status:       "succeeded",
			ApiUrl:       "/api/v1/builds/214953805",
			JobName:      "quickstart-smoke",
			PipelineId:   24,
			PipelineName: "concourse",
			StartTime:    time.Date(2023, time.October, 16, 14, 7, 45, 0, time.UTC),
			EndTime:      time.Date(2023, time.October, 16, 14, 11, 55, 0, time.UTC),
			CreatedBy:    "",
		},
	}
	qt.Assert(t, qt.DeepEquals(got, want))
}
