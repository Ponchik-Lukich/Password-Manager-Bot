package database

import (
	"context"
	"password-manager/cmd/models"
)

func AddUser(user models.User) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO users (chat_id, state) VALUES ($1, $2) ON CONFLICT DO NOTHING", user.ChatID, "wait")
	return err
}

func GetUser(userId int64) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE chat_id = $1`
	err := pool.QueryRow(context.Background(), query, userId).Scan(&user.ChatID, &user.State)
	return user, err
}

func SetUserState(chatID int64, state string) error {
	query := `UPDATE users SET state = $1 WHERE chat_id = $2`
	_, err := pool.Exec(context.Background(), query, state, chatID)
	return err
}
