package app

import (
	"github.com/kalunik/tgBot-urShorty/config"
	"github.com/kalunik/tgBot-urShorty/internal/bot"
	repo "github.com/kalunik/tgBot-urShorty/internal/repository"
	"github.com/kalunik/tgBot-urShorty/internal/usecase"
	"github.com/kalunik/tgBot-urShorty/internal/usecase/webapi"
	"github.com/kalunik/tgBot-urShorty/pkg/logger"
	bolt "go.etcd.io/bbolt"
)

type Bot struct {
	boltDB *bolt.DB
	log    logger.Logger
	conf   config.AppConfig
}

func NewBot(boltDB *bolt.DB, log logger.Logger, conf config.AppConfig) *Bot {
	return &Bot{
		boltDB: boltDB,
		log:    log,
		conf:   conf,
	}
}

func (b *Bot) Run() {
	botUsecase := usecase.NewCollectionUsecase(
		repo.NewBoltRepository(b.boltDB),
		webapi.NewShortenerWebAPI(b.conf))

	collectionBot, err := bot.NewCollectionBot(b.conf)
	if err != nil {
		b.log.Error(err)
		return
	}
	b.log.Infof("Authorized on account %s", collectionBot.Tgbot.Self.UserName)

	collectionHandlers := bot.NewCollectionBotHandlers(botUsecase, collectionBot.Tgbot)

	collectionBot.NewRouter(
		collectionHandlers,
	)
}
