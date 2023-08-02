package service

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/ipfs/go-cid"
	"github.com/switfs/switfs-block/utils/lotus-rpc"
)

func Event_Listening() error {
	//	 创建一个区块同步监听器
	ctx := context.Background()

	listener := make(chan []*api.HeadChange)
	var data map[cid.Cid]address.Address

	// 启动监听器
	go func() {
		for changes := range listener {
			log.Info("高度   》》》》》》》》》》》》 ", changes[0].Val.Height().String())
			data = make(map[cid.Cid]address.Address)
			for _, change := range changes {
				for _, block := range change.Val.Blocks() {
					log.Info("收到区块:", block.Cid().String(), "bk ", block.Miner.String())
					data[block.Cid()] = block.Miner
				}
			}
			for s, v := range data {
				fmt.Println(s, "   ", v)
			}
		}
	}()

	// 获取ChainSync API
	sub, err := lotus.Node.ChainNotify(ctx)
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
