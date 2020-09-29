package app

import "github.com/urfave/cli/v2"

type Command struct {
	commandlines []*cli.Command
}

func NewCommand() *Command {
	return &Command{
		commandlines: []*cli.Command{
			startCommand,
		},
	}
}

func (c *Command) Commandlines() []*cli.Command {
	return c.commandlines
}

var startCommand = &cli.Command{
	Name:  "start",
	Usage: "start a new gateway server",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "enable",
			Usage: "enable the controllers for the group",
		},
	},
	Action: func(clix *cli.Context) error {
		return nil
	},
}
