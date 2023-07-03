package main

import (
	"github.com/BurntSushi/toml"
	"github.com/filecoin-project/lotus/lib/lotuslog"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/cmd"
	"github.com/switfs/switfs-block/config"
	"github.com/switfs/switfs-block/utils/lotus-rpc"
	"github.com/switfs/switfs-block/utils/mysql-rpc"
	"github.com/switfs/switfs-block/utils/redis-rpc"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("main")

func init() {
	if _, err := toml.DecodeFile("./config.toml", &config.LotusConfig); err != nil {
		log.Error("配置文件初始失败", err.Error())
		return
	}
	if err := redis.New(); err != nil {
		log.Error("redis 初始失败 ", err.Error())
		return
	}
	if err := lotus.New(); err != nil {
		log.Error("lotus 初始化失败  ", err.Error())
		return
	}
	if err := mysql.New(); err != nil {
		log.Error("mysql 数据初始化失败 ", err.Error())
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
	mysql.RegisterTables()

	app.Setup()
	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}
