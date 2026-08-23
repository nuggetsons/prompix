package cli

import (
	"fmt"
	"time"

	"github.com/nuggetsons/prompix/internal/prometheus"
)

type RenderCmd struct {
	URL string `help:"Prometheus server URL." default:"http://localhost:9090"`

	Query string `help:"PromQL query to render." required:""`

	Range time.Duration `help:"Time range to query." default:"1h"`

	Step time.Duration `help:"Prometheus query step." default:"1m"`

	Output string `help:"Output image path." short:"o" default:"graph.png"`

	Width int `help:"Image width in pixels." short:"w" default:"1000"`

	Height int `help:"Image height in pixels." short:"h" default:"400"`
}

func (cmd *RenderCmd) Run() error {
	start := time.Now().Add(-cmd.Range).Format(time.RFC3339)
	end := time.Now().Format(time.RFC3339)

	step := fmt.Sprintf("%ds", int(cmd.Step.Seconds()))

	response, err := prometheus.Query(cmd.URL, cmd.Query, start, end, step)
	if err != nil {
		return fmt.Errorf("failed to query Prometheus: %w", err)
	}

	fmt.Printf("Prometheus Response: %+v\n", response)

	// call render

	return nil
}
