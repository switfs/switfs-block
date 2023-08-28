package cmd

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/switfs/switfs-block/service"
	"github.com/switfs/switfs-block/utils/lotus-rpc"
	"github.com/urfave/cli/v2"
	"os"
)

var RUN = &cli.Command{
	Name:  "run",
	Usage: "监控出块确认",
	Action: func(cctx *cli.Context) error {
		var addr []address.Address
		st, err := service.MinerId()
		if err != nil {
			return err
		}
		for _, v := range st {
			//fmt.Println(k, "===f", "f"+v)
			addres, _ := address.NewFromString("f" + v)
			addr = append(addr, addres)
		}

		head, err := lotus.Node.ChainHead(context.Background())
		if err != nil {
			return err
		}
		ts := head
		for _, v := range addr {
			count := 5
			for count > 0 {
				tsk := ts.Key()
				bhs := ts.Blocks()
				for _, bh := range bhs {
					if bh.Miner == v {
						fmt.Printf("%8d | %s | %s\n", ts.Height(), bh.Cid(), v.String())
						height := ts.Height() - 1
						err := service.MinerUP(bh.Cid().String(), height.String(), v.String())
						if err != nil {
							return err
						}
						count--
					} else {
						_, _ = fmt.Fprintf(os.Stderr, "\r\x1b[0KChecking epoch %s", cliutil.EpochTime(head.Height(), bh.Height))
					}
				}
				tsk = ts.Parents()
				ts, err = lotus.Node.ChainGetTipSet(context.Background(), tsk)
				if err != nil {
					return err
				}
			}

		}

		return nil
	},
}
