package app

import "github.com/urfave/cli"

func NewGateWayCommands() []cli.Command{
	return []cli.Command{
		startCommand,
	}
}

var startCommand = cli.Command{
	Name:  "start",
	Usage: "start a new gateway server",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "enable",
			Usage: "enable the controllers for the group",
		},
	},
	Action: func(clix *cli.Context) error {
		return nil
	},
}
