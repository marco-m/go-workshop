package loadmaster

import (
	"context"
	"fmt"

	"github.com/marco-m/go-workshop/loadmaster/pkg/concourse"
)

func CmdBuildTime(ctx context.Context, global Global, local *BuildTimeCmd) error {
	concourseClient, err := concourse.NewClient(concourse.Client{
		Server:     global.Server,
		HttpClient: global.HttpClient,
	})
	if err != nil {
		return fmt.Errorf("build-time: %s", err)
	}

	builds, err := concourseClient.ListPipelineBuilds(ctx, global.Team, local.Pipeline)
	if err != nil {
		return fmt.Errorf("build-time: %s", err)
	}

	for _, build := range builds {
		fmt.Println(build)
	}

	return nil
}
