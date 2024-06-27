package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/service"
	"github.com/switfs/switfs-block/utils/mysql-rpc"
	"github.com/urfave/cli/v2"

	"syscall"
)

const (
	VERSION = "v2.2.1"
)

var log = logging.Logger("main")

var RUN = &cli.Command{
	Name:  "run",
	Usage: "监控出块确认",
	Action: func(cctx *cli.Context) error {
		if len(os.Args) < 2 {
			printUsage()
			return nil
		}

		subCmd := os.Args[1]
		switch subCmd {

		case "run":
			sigCh := make(chan os.Signal, 2)
			signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
			service.Start()

			go createHttpServer()
			select {
			case sig := <-sigCh:
				log.Warn("received shutdown signal: ", sig)

			}
		default:
			printUsage()
		}

		db, _ := mysql.RPC.DB()

		db.Close()
		return nil
	},
}

func createHttpServer() {
	//logs.GetLogger().Info("release mode:", config.GetConfig().Release)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run("127.0.0.1:6530")
	if err != nil {
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Println("NAME:")
	fmt.Println("    switfs-block")
	fmt.Println("VERSION:")
	fmt.Println("    " + getVersion())
	fmt.Println("USAGE:")
	fmt.Println("    switfs-block version")
	fmt.Println("    switfs-block run")

}

func getVersion() string {
	return VERSION
}
