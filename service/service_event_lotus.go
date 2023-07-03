package service

import (
	"context"
	"github.com/filecoin-project/lotus/api"
	"github.com/switfs/switfs-block/utils/lotus-rpc"
)

func Event_Listening() error {
	//	 创建一个区块同步监听器
	ctx := context.Background()

	listener := make(chan []*api.HeadChange)

	// 启动监听器
	go func() {
		for changes := range listener {
			for _, change := range changes {
				for _, block := range change.Val.Blocks() {
					log.Info("收到区块:", block.Cid().String(), "bk ", block.Miner.String())
				}
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
				//case <-ctx.Done():
				//	return
			}
		}
	}()
	return nil
}
