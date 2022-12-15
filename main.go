package main

import (
	"github.com/FiiLabs/block_explorer/api"
	"github.com/FiiLabs/block_explorer/config"
	"github.com/FiiLabs/block_explorer/handlers"
	"github.com/FiiLabs/block_explorer/libs/logger"
	"github.com/FiiLabs/block_explorer/libs/pool"
	"github.com/FiiLabs/block_explorer/models"
	"github.com/gin-gonic/gin"
	"github.com/FiiLabs/block_explorer/tasks"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	c := make(chan os.Signal)

	defer func() {
		logger.Info("System Exit")

		models.Close()

		if err := recover(); err != nil {
			logger.Error("occur error", logger.Any("err", err))
			os.Exit(1)
		}
	}()

	conf, err := config.ReadConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}
	models.Init(conf)
	pool.Init(conf)
	handlers.InitRouter(conf)

	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	tasks.Start(tasks.NewSyncTask(conf))


	r := gin.Default()
	api.Routers(r)

	go func() {
		logrus.Fatal(r.Run(conf.Server.IpPort))
	}()

	<-c

	logrus.Println("Shutdown Server ...")
}
