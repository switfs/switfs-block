package main

import (
	"github.com/BurntSushi/toml"
	"github.com/filecoin-project/lotus/lib/lotuslog"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/cmd"
	"github.com/switfs/switfs-block/config"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("main")

func init() {
	if _, err := toml.DecodeFile("./config.toml", &config.LotusConfig); err != nil {
		log.Errorf("配置文件初始失败 %s", err.Error())
		return
	}
}

func main() {
	lotuslog.SetupLogLevels()
	log.Info("Start  ck  switfs  block 监听事件程序  .......")
	app := cli.App{
		Commands: []*cli.Command{
			cmd.Block,
			cmd.Run,
		},
	}

	app.Setup()
	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}
