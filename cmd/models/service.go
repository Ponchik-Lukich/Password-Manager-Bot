package models

type Service struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserID   int64  `json:"user_id"`
}
