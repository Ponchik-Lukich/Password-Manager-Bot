package database

import (
	"context"
	"password-manager/cmd/models"
)

func AddUser(user models.User) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO users (chat_id) VALUES ($1) ON CONFLICT DO NOTHING", user.ChatID)
	return err
}
