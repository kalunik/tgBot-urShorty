package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kalunik/tgBot-urShorty/internal/usecase"
	"strconv"
)

type CollectionBotHandlers interface {
	AddNewPath(chatID int64, fullUrl string, pathMeta string) (string, error)
	ListPaths(chatID int64) (map[string]string, error)
	GetPathVisits(chatID int64)
}

type collectionBotHandlers struct {
	collectionUC usecase.CollectionUsecase
	bot          *tgbotapi.BotAPI
}

func NewCollectionBotHandlers(uc usecase.CollectionUsecase, tgbot *tgbotapi.BotAPI) CollectionBotHandlers {
	return &collectionBotHandlers{collectionUC: uc, bot: tgbot}
}

func (h collectionBotHandlers) AddNewPath(chatID int64, fullUrl string, pathMeta string) (string, error) {
	collection, err := h.collectionUC.AddNewPath(strconv.FormatInt(chatID, 10), fullUrl, pathMeta)
	if err != nil {
		return "", err
	}
	return collection.MetaPath[pathMeta], nil
}

func (h collectionBotHandlers) ListPaths(chatID int64) (map[string]string, error) {
	list, err := h.collectionUC.ListPaths(strconv.FormatInt(chatID, 10))
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (h collectionBotHandlers) GetPathVisits(chatID int64) {
	//TODO implement me
	panic("implement me")
}
