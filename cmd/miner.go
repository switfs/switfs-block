package cmd

import (
	"fmt"
	"github.com/switfs/switfs-block/service"
	"github.com/urfave/cli/v2"
)

var MinerIdCmd = &cli.Command{
	Name:  "miner",
	Usage: "Manage miner id ",
	Subcommands: []*cli.Command{
		MinerAddCmd,
		MinerDelCmd,
		MinerListCmd,
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
		fmt.Println("添加成功.......", address)
		return nil
	},
}

var MinerDelCmd = &cli.Command{
	Name:  "del",
	Usage: "添加矿工miner Id",
	Action: func(cctx *cli.Context) error {
		address := cctx.Args().Get(0)
		if len(address) < 0 {
			log.Error("错误矿工id")
			return nil
		}
		mineradd := service.NewMinerIdService()
		if err := mineradd.Del(address); err != nil {
			log.Error(err)

			return err
		}
		fmt.Println("删除成功.......", address)
		return nil
	},
}

var MinerListCmd = &cli.Command{
	Name:  "list",
	Usage: "添加矿工miner Id",
	Action: func(cctx *cli.Context) error {

		mineradd := service.NewMinerIdService()
		data, err := mineradd.List()
		if err != nil {
			log.Error(err)
			return err
		}

		for k, v := range data {
			fmt.Println("k ", k, " MinerId ", v.MinerAddress)
		}

		return nil
	},
}
