package main

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("datadaotool")

func main() {
	//lotuslog.SetupLogLevels()

	local := []*cli.Command{
		getDealCmd,
		submitDealProposalCmd,
		//		eventsMonitorCmd,
	}

	app := &cli.App{
		Name:  "datadaotool",
		Usage: "",
		//Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "private-key",
				EnvVars: []string{"PRIVATE_KEY"},
				Value:   "",
			},
			&cli.StringFlag{
				Name: "endpoint",
				//Value: "ws://localhost:1234/rpc/v1",
				//Value: "ws://api.hyperspace.node.glif.io/rpc/v1",
				//Value: "wss://wss.hyperspace.node.glif.io/apigw/lotus",

				//Value: "wss://wss.hyperspace.node.glif.io/rpc/v1",
				//Value: "https://api.hyperspace.node.glif.io/rpc/v1",

				Value: "wss://wss.hyperspace.node.glif.io/apigw/lotus/rpc/v0",
			},
			&cli.StringFlag{
				Name:  "token",
				Value: "",
			},
		},

		Commands: local,
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		log.Errorf("%+v", err)
		os.Exit(1)
		return
	}
}
