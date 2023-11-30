package main

import (
	"github.com/kalunik/tgBot-urShorty/config"
	"github.com/kalunik/tgBot-urShorty/internal/app"
	"github.com/kalunik/tgBot-urShorty/pkg/db"
	"github.com/kalunik/tgBot-urShorty/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	log.InitLogger()
	log.Info("launching app")

	configDriver, err := config.LoadNewConfig()
	if err != nil {
		log.Fatal(err)
	}
	appConfig, err := configDriver.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	boltDB, err := db.NewBoltStorage(appConfig)

	app.NewBot(boltDB, log, appConfig).Run()

	if err = boltDB.Close(); err != nil {
		log.Error("boltDB close failure: ", err.Error())
	}
}
