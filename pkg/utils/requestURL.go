package utils

import (
	"fmt"
	"github.com/kalunik/tgBot-urShorty/config"
	"net/url"
	"strings"
)

func JoinReqURL(appConf config.AppConfig, apiRoutePath string) (string, error) {
	protocol := "http://"
	if appConf.URLShortenerAPI.SSL {
		protocol = "https://"
	}
	address := strings.Builder{}
	address.Write([]byte(appConf.URLShortenerAPI.ServerHost))
	address.Write([]byte(appConf.URLShortenerAPI.ServerPort))

	reqURL, err := url.JoinPath(
		protocol,
		address.String(),
		apiRoutePath)
	if err != nil {
		return "", fmt.Errorf("join path failure: %w", err)
	}
	return reqURL, nil
}
