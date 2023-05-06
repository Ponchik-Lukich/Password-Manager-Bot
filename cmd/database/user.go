package database

import (
	"context"
	"password-manager/cmd/models"
)

func AddUser(user models.User) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO users (chat_id, state) VALUES ($1, $2) ON CONFLICT DO NOTHING", user.ChatID, "wait")
	return err
}

func GetUser(userId int64, withMessage bool) (models.User, error) {
	var user models.User
	var err error
	if withMessage {
		query := `SELECT chat_id, state, message_id FROM users WHERE chat_id = $1`
		err = pool.QueryRow(context.Background(), query, userId).Scan(&user.ChatID, &user.State, &user.MessageID)
	} else {
		query := `SELECT chat_id, state FROM users WHERE chat_id = $1`
		err = pool.QueryRow(context.Background(), query, userId).Scan(&user.ChatID, &user.State)
	}
	return user, err
}

func SetUserState(chatID int64, state string, messageId ...int) error {
	if len(messageId) > 0 {
		query := `UPDATE users SET state = $1, message_id = $2 WHERE chat_id = $3`
		_, err := pool.Exec(context.Background(), query, state, messageId[0], chatID)
		return err
	} else {
		query := `UPDATE users SET state = $1 WHERE chat_id = $2`
		_, err := pool.Exec(context.Background(), query, state, chatID)
		return err
	}
}
