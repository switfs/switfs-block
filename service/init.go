package service

import (
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
	"sync"
)

var (
	log  = logging.Logger("service")
	lock sync.RWMutex
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
