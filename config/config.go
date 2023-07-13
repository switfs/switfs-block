package config

import (
	"github.com/BurntSushi/toml"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("config")

func init() {
	if _, err := toml.DecodeFile("./config.toml", &LotusConfig); err != nil {
		log.Errorf("配置文件初始失败 %s", err.Error())
		return
	}
}

var LotusConfig Config

type Config struct {
	Lotus *Lotus
	Mysql *Mysql
	Redis *Redis
}

type Lotus struct {
	Host  string
	Token string
}

type Mysql struct {
	Host   string
	User   string
	Pwd    string
	Dbname string
}

type Redis struct {
	Host   string
	DbName int
}
