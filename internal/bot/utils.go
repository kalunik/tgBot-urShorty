package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kalunik/tgBot-urShorty/pkg/utils"
)

func (c *CollectionBot) setBotAndMenuLanguage(languageCode string) {
	if languageCode == "" || c.langCode == languageCode {
		return
	}
	c.langCode = languageCode
	c.messagePack = utils.NewMessagePack(languageCode)

	c.setMenuButton()
}

func (c *CollectionBot) setMenuButton() {
	_, _ = c.Tgbot.Send(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "newlink",
			Description: c.messagePack.MenuNewLink()},
		tgbotapi.BotCommand{
			Command:     "mylinks",
			Description: c.messagePack.MenuMyLinks(),
		}))
}

func (c *CollectionBot) sendMessage(chatID int64, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	message, err := c.Tgbot.Send(msg)
	if err != nil {
		return tgbotapi.Message{}, fmt.Errorf("send message failure, %w", err)
	}
	return message, nil
}
