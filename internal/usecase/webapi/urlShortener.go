package webapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kalunik/tgBot-urShorty/config"
	"github.com/kalunik/tgBot-urShorty/internal/entity"
	"github.com/kalunik/tgBot-urShorty/pkg/utils"
	"net/http"
)

type UrlShortenerWebAPI struct {
	config config.AppConfig
}

func NewShortenerWebAPI(appConfig config.AppConfig) *UrlShortenerWebAPI {
	return &UrlShortenerWebAPI{config: appConfig}
}

func (w UrlShortenerWebAPI) AddPath(fullUrl string) (string, error) {
	postBody, err := json.Marshal(map[string]string{
		"full_url": fullUrl,
	})
	if err != nil {
		return "", errors.New("request body marshalling failure")
	}
	reqBody := bytes.NewBuffer(postBody)
	reqURL, err := utils.JoinReqURL(w.config, "/shorten/")
	if err != nil {
		return "", fmt.Errorf("joinReqURL: %w", err)
	}

	resp, err := http.Post(reqURL, "application/json", reqBody)
	if err != nil {
		return "", fmt.Errorf("post request failure: %w", err)
	}
	defer resp.Body.Close()

	responseData := &entity.ResponseAddPath{}

	if err = json.NewDecoder(resp.Body).Decode(responseData); err != nil {
		return "", errors.New("decode failure")
	}
	return responseData.Short, nil
}

func (w UrlShortenerWebAPI) name() {

}
