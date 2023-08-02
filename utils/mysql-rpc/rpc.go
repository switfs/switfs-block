package mysql

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/config"
	models2 "github.com/switfs/switfs-block/server/models"
	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	log = logging.Logger("mysql")
	RPC *gorm.DB
	err error
)

func InitNew() {
	if err := New(); err != nil {
		log.Error(err.Error())
		return
	}

	err = RPC.AutoMigrate(
		new(models2.Miner),
		new(models2.BlockTotal),
	)

	if err != nil {
		log.Error(err.Error())
		return
	}
}

func New() error {
	dsn := config.LotusConfig.Mysql.User + ":" + config.LotusConfig.Mysql.Pwd + "@(" + config.LotusConfig.Mysql.Host + ")/" + config.LotusConfig.Mysql.Dbname + "?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci&readTimeout=10s&writeTimeout=10s"
	RPC, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Errorf("[db connection failed] Database name: wallet_address %v", err)
		return xerrors.Errorf("[db connection failed] Database name: wallet_address %w", err)
	}

	sqlDB, err := RPC.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(128)
	sqlDB.SetConnMaxLifetime(120 * time.Second)
	sqlDB.SetConnMaxIdleTime(60 * time.Second)

	return nil
}
