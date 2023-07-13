package lotus

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/v1api"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/config"
	"net/http"
)

var log = logging.Logger("lotus")

var (
	Node v1api.FullNodeStruct
)

func init() {
	if err := New(); err != nil {
		log.Errorf("lotus cconnecting do %s ", err.Error())
		return
	}
}

func New() error {
	headers := http.Header{"Authorization": []string{"Bearer " + config.LotusConfig.Lotus.Token}}
	_, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+config.LotusConfig.Lotus.Host+"/rpc/v0", "Filecoin", []interface{}{&Node.Internal, &Node.CommonStruct.Internal}, headers)
	if err != nil {
		log.Errorf("connecting with lotus failed: %s", err)
		return err
	}
	return nil
}
