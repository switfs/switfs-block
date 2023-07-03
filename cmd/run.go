package cmd

import (
	lcli "github.com/filecoin-project/lotus/cli"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/service"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

var (
	g   errgroup.Group
	log = logging.Logger("cmd")
)
var Run = &cli.Command{
	Name:  "run",
	Usage: "start sync check",
	Action: func(ctxx *cli.Context) error {

		log.Info("start sss 1")
		chainAPI, ncloser, err := lcli.GetFullNodeAPIV1(ctxx)
		if err != nil {
			return xerrors.Errorf("getting full node api: %w", err)
		}
		defer ncloser()
		ctx := lcli.ReqContext(ctxx)
		err = service.Event_Listening(ctx, chainAPI)
		if err != nil {
			return err
		}
		server := &http.Server{
			Addr:         "127.0.0.1",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		g.Go(func() error {
			return server.ListenAndServe()
		})
		if err := g.Wait(); err != nil {
			log.Error(err.Error())
		}

		return nil
	},
}
