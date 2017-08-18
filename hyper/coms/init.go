package coms

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"github.com/wallnutkraken/hyper/hservice"
)

func Init(c *cli.Context) {
	dae, err := hservice.New()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	if err := dae.Install(); err != nil {
		fmt.Printf("init: %s\n", err.Error())
		os.Exit(1)
	}
}

