package coms

import (
	"github.com/urfave/cli"
	"github.com/wallnutkraken/hyper/hservice"
	"fmt"
	"os"
)

func Remove(c *cli.Context) {
	dae, err := hservice.New()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	if err := dae.Uninstall(); err != nil {
		fmt.Printf("remove: %s\n", err.Error())
		os.Exit(1)
	}
}