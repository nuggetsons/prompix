package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nuggetsons/prompix/internal/prometheus"
	"github.com/nuggetsons/prompix/internal/render"
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

	status := response.Status
	if status != "success" {
		return fmt.Errorf("Prometheus query failed with status: %s", status)
	}

	results := response.Data.Result
	all_timeseries := make(map[string][][]float64)

	for _, r := range results {
		metric := r.Metric
		values := r.Values

		labels := ""

		for key, val := range metric {
			labels += key + "=" + val + ","
		}

		labels = labels[:len(labels)-1]

		timeseries, err := cmd.convert(values)
		if err != nil {
			return err
		}

		all_timeseries[labels] = timeseries
	}

	err = render.Render(cmd.Query, all_timeseries, cmd.Width, cmd.Height, cmd.Output)
	if err != nil {
		return err
	}

	fmt.Printf("Saved to %s\n", cmd.Output)
	return nil
}

func (cmd *RenderCmd) convert(values []any) ([][]float64, error) {
	result := make([][]float64, len(values))

	for i := range values {
		point, ok := values[i].([]any)
		if !ok {
			return [][]float64{}, fmt.Errorf("failed to convert value to []any")
		}

		time, ok := point[0].(float64)
		if !ok {
			return [][]float64{}, fmt.Errorf("failed to convert time to float64")
		}

		valStr, ok := point[1].(string)
		if !ok {
			return [][]float64{}, fmt.Errorf("failed to convert value to string")
		}

		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return [][]float64{}, fmt.Errorf("failed to parse value as float64: %w", err)
		}

		result[i] = []float64{time, val}
	}

	return result, nil
}
