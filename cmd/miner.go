package cmd

import (
	"github.com/urfave/cli/v2"
)

var Miner = &cli.Command{
	Name:  "add",
	Usage: "添加矿工miner Id",
	Action: func(cctx *cli.Context) error {
		address := cctx.Args().Get(0)
		if len(address) < 0 {
			log.Error("错误矿工id")
			return nil
		}
		//addes := cctx.Args().Get(0)

		return nil
	},
}
