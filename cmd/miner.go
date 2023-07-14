package cmd

import (
	"github.com/switfs/switfs-block/service"
	"github.com/urfave/cli/v2"
)

var MinerIdCmd = &cli.Command{
	Name:  "miner",
	Usage: "Manage miner id ",
	Subcommands: []*cli.Command{
		MinerAddCmd,
	},
}

var MinerAddCmd = &cli.Command{
	Name:  "add",
	Usage: "添加矿工miner Id",
	Action: func(cctx *cli.Context) error {
		address := cctx.Args().Get(0)
		if len(address) < 0 {
			log.Error("错误矿工id")
			return nil
		}
		mineradd := service.NewMinerIdService()
		if err := mineradd.Add(address); err != nil {
			log.Error(err)
			return err
		}

		return nil
	},
}
