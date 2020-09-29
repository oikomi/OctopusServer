package main

import (
	"github/oikomi/OctopusServer/server/cmd/gateway/app"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	gateway := cli.NewApp()
	gateway.Name = "gateway"
	gateway.Version = "1"
	gateway.Usage = "gateway server"
	gateway.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output in the logs",
		},
	}
	gateway.Commands = app.NewCommand().Commandlines()

	gateway.Before = func(clix *cli.Context) error {
		if clix.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	if err := gateway.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
