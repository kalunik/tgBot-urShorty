package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kalunik/tgBot-urShorty/config"
	"github.com/kalunik/tgBot-urShorty/pkg/logger"
	"github.com/kalunik/tgBot-urShorty/pkg/utils"
)

type CollectionBot struct {
	Tgbot         *tgbotapi.BotAPI
	langCode      string
	messagePack   utils.Messages
	log           logger.Logger
	webAPIAddress string
}

func NewCollectionBot(conf config.AppConfig) (*CollectionBot, error) {
	bot, err := tgbotapi.NewBotAPI(conf.TelegramAPI.Token)
	if err != nil {
		return nil, fmt.Errorf("newBotAPI failure: %w", err)
	}
	bot.Debug = true

	address, err := utils.JoinReqURL(conf, "/")
	if err != nil {

	}
	return &CollectionBot{Tgbot: bot, webAPIAddress: address}, nil
}

func (c *CollectionBot) NewRouter(collHandle CollectionBotHandlers) {
	var (
		addCommandTrigger    uint8
		addCommandPrevAnswer string
	)

	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60
	u.AllowedUpdates = []string{
		tgbotapi.UpdateTypeMessage,
		tgbotapi.UpdateTypeCallbackQuery,
	}
	updates := c.Tgbot.GetUpdatesChan(u)
	updates.Clear()

	c.setBotAndMenuLanguage("en")

	for update := range updates {

		if update.SentFrom() != nil {
			c.setBotAndMenuLanguage(update.SentFrom().LanguageCode)
		}

		switch {
		case update.CallbackQuery != nil:
			addCommandTrigger = 0
			c.callbackMux(update.CallbackQuery, collHandle)
		case update.Message.IsCommand():
			addCommandTrigger = 0
			c.commandMux(&addCommandTrigger, update.Message, collHandle)
		case addCommandTrigger > 0 && update.Message != nil:
			addCommandTrigger--
			c.commandAnswer(update, collHandle, &addCommandPrevAnswer)
		default:
			c.sendMessage(update.Message.Chat.ID, c.messagePack.Default())
		}
	}
}
