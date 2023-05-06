package models

type Service struct {
	ID         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Login      string `json:"login,omitempty"`
	Password   string `json:"password,omitempty"`
	UserChatID int64  `json:"user_id,omitempty"`
}
