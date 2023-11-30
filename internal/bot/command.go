package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kalunik/tgBot-urShorty/pkg/utils"
	"reflect"
)

func (c *CollectionBot) commandMux(addCommandTrigger *uint8, message *tgbotapi.Message, collHandle CollectionBotHandlers) {
	c.messagePack = utils.NewMessagePack(message.From.LanguageCode)

	switch message.Command() {
	case "start":
		c.sendMessage(message.Chat.ID, c.messagePack.Start())
	case "help":
		c.sendMessage(message.Chat.ID, c.messagePack.Help())
	case "newlink":
		c.sendMessage(message.Chat.ID, c.messagePack.NewLink())
		*addCommandTrigger = 2
	case "mylinks":
		c.myLinksCommand(message, collHandle)
	default:
		c.sendMessage(message.Chat.ID, c.messagePack.Default())
	}
}

func (c *CollectionBot) commandAnswer(update tgbotapi.Update, collHandle CollectionBotHandlers, fullUrl *string) {
	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String &&
		update.Message.Text != "" {

		if *fullUrl == "" {
			*fullUrl = update.Message.Text
			c.sendMessage(update.Message.Chat.ID, c.messagePack.AddLinkNote())
			return
		}
		shortUrl, err := collHandle.AddNewPath(update.Message.Chat.ID, *fullUrl, update.Message.Text)
		*fullUrl = ""
		if err != nil {
			c.sendMessage(update.Message.Chat.ID, err.Error())
			return
		}

		c.sendMessage(update.Message.Chat.ID, shortUrl)
	}
}

func (c *CollectionBot) myLinksCommand(message *tgbotapi.Message, collHandle CollectionBotHandlers) {
	MetaPath, err := collHandle.ListPaths(message.Chat.ID)
	if err != nil {
		c.sendMessage(message.Chat.ID, err.Error())
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, c.messagePack.MyLinks())
	urlList := tgbotapi.NewInlineKeyboardMarkup()
	for k, v := range MetaPath {
		button := tgbotapi.NewInlineKeyboardButtonData(k, utils.MergeString(k, v))
		urlList.InlineKeyboard = append(urlList.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(button))
	}
	msg.ReplyMarkup = urlList
	c.Tgbot.Send(msg)
}
