package cmd

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/service"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

		sigCh := make(chan os.Signal, 1)
		err := service.Event_Listening()
		if err != nil {
			return err
		}
		server := &http.Server{
			Addr:         "127.0.0.1:9871",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		g.Go(func() error {
			return server.ListenAndServe()
		})

		if err := g.Wait(); err != nil {
			log.Error(err.Error())
		}
		<-sigCh

		return nil
	},
}

func BlcokRoute() {

}
