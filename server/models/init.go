package models

import (
	logging "github.com/ipfs/go-log/v2"
	"sync"
)

var (
	log  = logging.Logger("modles")
	lock sync.RWMutex
)
