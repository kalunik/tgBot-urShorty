package usecase

import (
	"errors"
	"github.com/kalunik/tgBot-urShorty/internal/entity"
	"github.com/kalunik/tgBot-urShorty/internal/repository"
	"github.com/kalunik/tgBot-urShorty/internal/usecase/webapi"
)

type CollectionUsecase interface {
	AddNewPath(chatID string, fullUrl string, pathNote string) (entity.Collection, error)
	ListPaths(chatID string) (map[string]string, error)
	GetPathVisits() error
}

type CollectionUC struct {
	boltRepo     repository.BoltRepository
	shortenerAPI *webapi.UrlShortenerWebAPI
}

func NewCollectionUsecase(repo repository.BoltRepository, api *webapi.UrlShortenerWebAPI) CollectionUsecase {
	return &CollectionUC{boltRepo: repo, shortenerAPI: api}
}

func (u CollectionUC) AddNewPath(chatID string, fullUrl string, pathMeta string) (entity.Collection, error) {

	path, err := u.shortenerAPI.AddPath(fullUrl)
	if err != nil {
		return entity.Collection{}, errors.New("request to server fail, try again later")
	}

	collection, err := u.boltRepo.Get(chatID)
	var emptyDb *repository.EmptyDatabaseError

	if errors.As(err, &emptyDb) {
		err = nil
		collection = entity.Collection{
			ChatID:   chatID,
			MetaPath: make(map[string]string),
		}
	} else if err != nil {
		return entity.Collection{}, err
	}
	collection.MetaPath[pathMeta] = path

	if err = u.boltRepo.Set(collection); err != nil {
		return entity.Collection{}, err
	}

	return collection, nil
}

func (u CollectionUC) ListPaths(chatID string) (map[string]string, error) {
	collection, err := u.boltRepo.Get(chatID)
	if err != nil {
		return nil, err
	}
	return collection.MetaPath, nil
}

func (u CollectionUC) GetPathVisits() error {
	return nil
}
