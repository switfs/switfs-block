package cmd

import (
	"github.com/switfs/switfs-block/service"
	"github.com/urfave/cli/v2"
)

var RUN = &cli.Command{
	Name:  "run",
	Usage: "监控出块确认",
	Action: func(cctx *cli.Context) error {

		st, err := service.MinerId()
		if err != nil {
			return err
		}
		for _, v := range st {
			//fmt.Println(k, "===f", "f"+v)
			service.Getdata("f" + v)

		}

		return nil
	},
}
