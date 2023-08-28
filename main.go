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
		log.Errorf("配置文件初始失败 %s", err.Error())
		return
	}

	if err := lotus.New(); err != nil {
		log.Errorf("lotus cconnecting do %s ", err.Error())
		return
	}

	mysql.InitNew()

	if err := redis.New(); err != nil {
		log.Errorf("redist error %s", err.Error())
		return
	}

}

func main() {
	lotuslog.SetupLogLevels()
	app := cli.App{
		Commands: []*cli.Command{
			cmd.RUN,
		},
	}

	app.Setup()
	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}
