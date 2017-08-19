package coms

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"github.com/wallnutkraken/hyper/hservice"
)

func Init(c *cli.Context) {
	if !hservice.IsSupported() {
		fmt.Println("error: platform not supported")
		os.Exit(1)
	}
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

