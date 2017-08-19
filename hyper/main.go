package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/wallnutkraken/hyper/hyper/coms"
)

func main() {
	/* Initialize the cli */
	hyperApp := cli.NewApp()
	hyperApp.Name = "hyper"
	hyperApp.Usage = "hyper [command] [arguments]"
	hyperApp.Version = version

	/* Set up the commands */
	hyperApp.Commands = []cli.Command {
		cli.Command{
			Name: "start",
			Usage: "Used to start the hyper daemon",
			Action: coms.Start,
		},
		cli.Command{
			Name: "init",
			Usage: "Installs the daemon into the operating system",
			Action: coms.Init,
		},
		cli.Command{
			Name: "remove",
			Usage: "Removes the daemon from the system",
			Action: coms.Remove,
		},
	}

	hyperApp.Run(os.Args)
}