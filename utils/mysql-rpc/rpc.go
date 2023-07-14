package mysql

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/config"
	"github.com/switfs/switfs-block/server/model"
	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	log = logging.Logger("mysql")
	RPC *gorm.DB
	err error
)

type TGormDataSourceManager struct {
	Models []interface{}
}

func init() {

	d := []interface{}{
		new(model.BlockTotal),
		new(model.Miner),
	}

	if err = New(); err != nil {
		log.Error(err.Error())
		return
	}

	models := TGormDataSourceManager{}
	models.RegisterModels(d...)
}

func New() error {
	var st = &TGormDataSourceManager{}

	dsn := config.LotusConfig.Mysql.User + ":" + config.LotusConfig.Mysql.Pwd + "@(" + config.LotusConfig.Mysql.Host + ")/" + config.LotusConfig.Mysql.Dbname + "?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci&readTimeout=10s&writeTimeout=10s"
	RPC, err = gorm.Open("mysql", mysql.Open(dsn))
	if err != nil {
		log.Errorf("[db connection failed] Database name: wallet_address %v", err)
		return xerrors.Errorf("[db connection failed] Database name: wallet_address %w", err)
	}

	RPC.DB().SetMaxOpenConns(100)
	RPC.DB().SetMaxIdleConns(128)
	RPC.DB().SetConnMaxLifetime(120 * time.Second)
	RPC.DB().SetConnMaxIdleTime(60 * time.Second)
	RPC.AutoMigrate(st.Models...)
	return nil
}

// auto create/migrate table if you want
func (manager *TGormDataSourceManager) RegisterModels(models ...interface{}) {
	manager.Models = append(manager.Models, models...)
}
