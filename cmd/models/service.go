package models

type Service struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	UserChatID int64  `json:"user_id"`
}
