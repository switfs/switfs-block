package main

import (
	"os"

	"github.com/BurntSushi/toml"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/cmd"
	"github.com/switfs/switfs-block/config"
	"github.com/switfs/switfs-block/utils/mysql-rpc"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("main")

func init() {
	if _, err := toml.DecodeFile("./config.toml", &config.LotusConfig); err != nil {
		log.Errorf("配置文件初始失败 %s", err.Error())
		return
	}

	mysql.InitNew()

}

func main() {
	_ = logging.SetLogLevel("*", "INFO")
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
