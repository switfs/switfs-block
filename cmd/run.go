package cmd

import (
	"context"
	"fmt"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v1api"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	"net/http"
)

var log = logging.Logger("cmd")
var Run = &cli.Command{
	Name:  "run",
	Usage: "start sync check",
	Action: func(ctxx *cli.Context) error {
		r := gin.Default()
		log.Info("start sss 1")
		chainAPI, ncloser, err := lcli.GetFullNodeAPIV1(ctxx)
		if err != nil {
			return xerrors.Errorf("getting full node api: %w", err)
		}
		defer ncloser()
		ctx := lcli.ReqContext(ctxx)
		log.Info("start sss 23")
		err = StartBlock(ctx, chainAPI)
		if err != nil {
			return err
		}
		log.Info("start sss4")

		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		r.Run()

		return nil
	},
}

func StartBlock(ctx context.Context, chainAPI v1api.FullNode) error {
	log.Info("start sss 2")
	//	 创建一个区块同步监听器
	listener := make(chan []*api.HeadChange)

	// 启动监听器
	go func() {
		for changes := range listener {
			for _, change := range changes {
				for _, block := range change.Val.Blocks() {
					fmt.Println("收到区块:", block.Cid().String(), "bk ", block.Miner.String())
				}
			}
		}
	}()
	// 获取ChainSync API
	sub, err := chainAPI.ChainNotify(ctx)
	if err != nil {
		panic(err)
	}
	// 开始监听区块同步事件
	go func() {
		for {
			select {
			case changes := <-sub:
				listener <- changes
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}
