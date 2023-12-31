package dto

import (
	logging "github.com/ipfs/go-log/v2"
	"sync"
)

var (
	log  = logging.Logger("dto")
	lock sync.RWMutex
)
