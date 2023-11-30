package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kalunik/tgBot-urShorty/pkg/utils"
	"net/url"
)

func (c *CollectionBot) callbackMux(query *tgbotapi.CallbackQuery, collHandle CollectionBotHandlers) {
	switch query.Data {
	case "":

	default:
		c.myLinksCallback(query, collHandle)
	}

}

func (c *CollectionBot) myLinksCallback(query *tgbotapi.CallbackQuery, collHandle CollectionBotHandlers) {
	MetaPath := utils.UnmergeStrings(query.Data)

	shortLink, _ := url.JoinPath(c.webAPIAddress, MetaPath[1])

	msg := tgbotapi.NewMessage(query.Message.Chat.ID, c.messagePack.ManageLink(MetaPath[0], shortLink))
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Visits", "Visits"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Delete", "Delete"),
			tgbotapi.NewInlineKeyboardButtonData("Edit note", "Edit note"),
		),
	)
	_, err := c.Tgbot.Send(msg)
	if err != nil {
		c.sendMessage(query.Message.Chat.ID, c.messagePack.InternalError())
	}
}
