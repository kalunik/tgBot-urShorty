package entity

type Collection struct {
	ChatID   string            `json:"chat_id"`
	MetaPath map[string]string `json:"meta_path"`
}
