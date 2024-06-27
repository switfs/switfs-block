package config

var LotusConfig Config

type Config struct {
	Mysql *Mysql
}

type Mysql struct {
	Host   string
	User   string
	Pwd    string
	Dbname string
}
