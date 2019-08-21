package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dynastymasra/avalon/infrastructure/provider"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/dynastymasra/avalon/infrastructure/web"

	"github.com/dynastymasra/avalon/infrastructure/provider/postgres"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/dynastymasra/avalon/config"
)

func init() {
	config.Load()
	config.SetupLogger()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Database initialization
	postgresDB, err := postgres.Connect()
	if err != nil {
		logrus.WithError(err).Fatalln("Unable to open connection to postgres")
	}

	providerInstance := provider.NewInstance(postgresDB)

	clientApp := cli.NewApp()
	clientApp.Name = config.ServiceName
	clientApp.Version = config.Version

	clientApp.Action = func(c *cli.Context) error {
		server := &graceful.Server{
			Timeout: 0,
		}
		go web.Run(server, providerInstance)
		select {
		case sig := <-stop:
			<-server.StopChan()
			logrus.Warningln(fmt.Sprintf("Service shutdown because %+v", sig))

			if err := postgresDB.Close(); err != nil {
				logrus.WithError(err).Errorln("Unable to turn off Postgres connections")
			}
			logrus.Infoln("Postgres Connection closed")

			os.Exit(0)
		}
		return nil
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
