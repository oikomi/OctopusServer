package main

import (
	"fmt"
	"github.com/oikomi/OctopusServer/server/cmd/gateway/app"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"math/rand"
	"os"
	"time"
)

const GATE_WAY_SERVER_CONF_FILE = "./gateway.yaml"

func main() {
	rand.Seed(time.Now().UnixNano())

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug, d",
			Usage:   "enable debug output in the logs",
			EnvVars: []string{"ENABLE_DEBUG"},
			FilePath: "GATE_WAY_SERVER_CONF_FILE",
		},
	}

	gateway := cli.App{
		Name:        "gateway",
		HelpName:    "gateway",
		Version:     "1",
		Description: "gateway server",
		Flags: flags,
		Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc(GATE_WAY_SERVER_CONF_FILE)),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "harold.miao",
				Email: "miaohong@miaohong.org",
			},
		},
		Copyright: "",
		Commands:  app.NewCommand().Commandlines(),
	}

	gateway.Before = func(clix *cli.Context) error {
		if clix.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	if err := gateway.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
