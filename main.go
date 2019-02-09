package main

import (
	"os"

	"github.com/giuem/ga-proxy/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ga-proxy"
	app.HideVersion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "ip, i",
			Value:  "127.0.0.1",
			Usage:  "`IP` to listen",
			EnvVar: "IP",
		},
		cli.StringFlag{
			Name:   "port, p",
			Value:  "9080",
			Usage:  "`port` to listen",
			EnvVar: "PORT",
		},
	}

	app.Action = func(c *cli.Context) error {
		server.Run(c.String("ip"), c.String("port"))
		return nil
	}

	app.Run(os.Args)
}
