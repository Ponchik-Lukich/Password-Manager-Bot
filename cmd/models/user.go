package models

type User struct {
	ChatID int64  `json:"chat_id"`
	State  string `json:"state"`
}
