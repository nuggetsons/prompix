package cli

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	Render RenderCmd `cmd:"" help:"Render a Prometheus query as an image."`
}

func Execute() error {
	var cli CLI

	ctx := kong.Parse(
		&cli,
		kong.Name("prompix"),
		kong.Description("Render Prometheus queries as images."),
	)

	return ctx.Run()
}