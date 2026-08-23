package main

import (
	"os"

	"github.com/nuggetsons/prompix/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}