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

func main() {
	rand.Seed(time.Now().UnixNano())

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "conf_file",
			Aliases: []string{"c"},
			Usage:   "gateway conf file",
			EnvVars: []string{"GATEWAY_CONF_FILE"},
		},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "enable debug output in the logs",
			EnvVars: []string{"ENABLE_DEBUG"},
		}),
	}

	gateway := cli.App{
		Name:        "gateway",
		Version:     "1",
		Description: "gateway server",
		Flags:       flags,
		Authors: []*cli.Author{
			{
				Name:  "harold.miao",
				Email: "miaohong@miaohong.org",
			},
		},
		Copyright: "",
		Commands:  app.NewCommand().Commandlines(),
	}

	gateway.Before = func(clix *cli.Context) error {
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
		if clix.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		logrus.Info("load conf file")
		confLoadFunc := altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("conf_file"))
		err := confLoadFunc(clix)
		if err != nil {
			logrus.WithError(err).Errorf("[gateway] failed to load conf file %s", clix.String("conf_file"))
			panic(err)
		}
		return nil
	}

	gateway.Action = func(ctx *cli.Context) error {
		debug := ctx.Bool("debug")
		fmt.Printf("debug is %v\n", debug)
		return nil
	}

	if err := gateway.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
