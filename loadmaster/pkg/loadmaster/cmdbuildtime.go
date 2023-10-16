package loadmaster

import (
	"context"
	"fmt"
	"os"
	"slices"
	"text/tabwriter"
	"time"

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

	type stats struct {
		JobName  string
		Count    int
		Duration time.Duration
	}

	// jobName -> stats
	JobStatsMap := make(map[string]*stats, 100)

	for _, build := range builds {
		if build.Status == concourse.StatusStarted ||
			build.Status == concourse.StatusPending {
			continue
		}
		d := build.EndTime.Sub(build.StartTime)
		if d < 0 {
			fmt.Println("negative time!!!")
			continue
		}
		if _, ok := JobStatsMap[build.JobName]; !ok {
			JobStatsMap[build.JobName] = &stats{JobName: build.JobName}
		}
		JobStatsMap[build.JobName].Count++
		JobStatsMap[build.JobName].Duration += d
	}

	JobStats := make([]*stats, 0, len(JobStatsMap))
	for _, v := range JobStatsMap {
		JobStats = append(JobStats, v)
	}

	fmt.Println("cumulative job build durations for pipeline", local.Pipeline)

	slices.SortFunc(JobStats, func(a, b *stats) int {
		if a.Duration < b.Duration {
			return 1
		}
		return -1
	})

	w := new(tabwriter.Writer)
	minwidth := 5
	tabwidth := 0
	padding := 2
	padchar := byte(' ')
	flags := uint(0)
	w.Init(os.Stdout, minwidth, tabwidth, padding, padchar, flags)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "job", "count", "average", "total")
	for _, j := range JobStats {
		count := fmt.Sprintf("%5d", j.Count)
		average := time.Duration(int(j.Duration) / j.Count).Round(time.Second)
		fmt.Fprintf(w, "%s\t%s\t%v\t%v\n", j.JobName, count, average, j.Duration)
	}
	w.Flush()

	return nil
}
