package utils

import "fmt"

type Messages interface {
	BadRequest() string
	InternalError() string
	Start() string
	Help() string
	Default() string
	NewLink() string
	AddLinkNote() string
	MyLinks() string
	MenuNewLink() string
	MenuMyLinks() string
	ManageLink(note string, link string) string
}

type LanguagePack struct {
	languageCode string
}

func NewMessagePack(languageCode string) Messages {
	return &LanguagePack{
		languageCode: supportedLanguage(languageCode),
	}
}

func supportedLanguage(languageCode string) string {
	supportedLanguages := []string{"en", "ru"}
	for _, lang := range supportedLanguages {
		if languageCode == lang {
			return languageCode
		}
	}
	return "en"
}

func (l LanguagePack) useLocale() MessageSet {
	switch l.languageCode {
	case "en":
		return enLocaleMessages()
	case "ru":
		return ruLocaleMessages()
	}
	return enLocaleMessages()
}

func (l LanguagePack) BadRequest() string {
	return l.useLocale()["BadRequest"]
}

func (l LanguagePack) InternalError() string {
	return l.useLocale()["InternalError"]
}

func (l LanguagePack) Start() string {
	return l.useLocale()["Start"]
}

func (l LanguagePack) Help() string {
	return l.useLocale()["Help"]
}

func (l LanguagePack) Default() string {
	return l.useLocale()["Default"]
}

func (l LanguagePack) NewLink() string {
	return l.useLocale()["NewLink"]
}

func (l LanguagePack) AddLinkNote() string {
	return l.useLocale()["AddLinkNote"]
}

func (l LanguagePack) MyLinks() string {
	return l.useLocale()["MyLinks"]
}

func (l LanguagePack) MenuNewLink() string {
	return l.useLocale()["menuNewLink"]
}

func (l LanguagePack) MenuMyLinks() string {
	return l.useLocale()["menuMyLinks"]
}

func (l LanguagePack) ManageLink(note string, link string) string {
	return fmt.Sprintf(l.useLocale()["ManageLink"], note, link)
}
