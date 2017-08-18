package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	/* Initialize the cli */
	hyperApp := cli.NewApp()
	hyperApp.Name = "hyper"
	hyperApp.Usage = "hyper [command] [arguments]"
	hyperApp.Version = version

	/* Set up the commands */
	hyperApp.Commands = []cli.Command {
	}

	hyperApp.Run(os.Args)
}