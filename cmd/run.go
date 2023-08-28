package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
	cors "github.com/itsjamie/gin-cors"
	"github.com/switfs/switfs-block/service"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"

	"syscall"
	"time"
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

		return nil
	},
}

func createHttpServer() {
	//logs.GetLogger().Info("release mode:", config.GetConfig().Release)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

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
