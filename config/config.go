package config

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("config")

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
