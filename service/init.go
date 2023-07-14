package service

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
	"sync"
)

var (
	log = logging.Logger("service")
	l   sync.Mutex
)

func closeTx(tx *gorm.DB, err *error) {
	r := recover()
	if r != nil {
		tx.Rollback()
		log.Error(r)
		*err = xerrors.New("panic")
		return
	}

	if *err != nil {
		tx.Rollback()
		log.Errorf("%+v", *err)
		return
	}
	tx.Commit()
}
