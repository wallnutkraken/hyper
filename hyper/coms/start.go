package coms

import (
	"github.com/urfave/cli"
	"github.com/wallnutkraken/hyper/hservice"
	"fmt"
	"os"
)

func Start(c *cli.Context) {
	dae, err := hservice.New()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	if err := dae.Start(); err != nil {
		fmt.Printf("start: %s\n", err.Error())
		os.Exit(1)
	}
}
