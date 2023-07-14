package main

import (
	"github.com/filecoin-project/lotus/lib/lotuslog"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("main")

func main() {
	lotuslog.SetupLogLevels()
	app := cli.App{
		Commands: []*cli.Command{
			cmd.Block,
			cmd.Run,
			cmd.MinerIdCmd,
		},
	}

	app.Setup()
	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}
