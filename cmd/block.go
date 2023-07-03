package cmd

import (
	"github.com/urfave/cli/v2"
)

var Block = &cli.Command{
	Name:  "block",
	Usage: "统计出块",
	Action: func(cctx *cli.Context) error {

		return nil
	},
}
