package database

import (
	"context"
	"errors"
	"log"
	"password-manager/cmd/models"
)

func AddService(service models.Service) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO services (name, login, password, user_chat_id)"+
		"VALUES ($1, $2, $3, $4)"+
		"ON CONFLICT (name, user_chat_id)"+
		"DO UPDATE SET"+
		"  login = EXCLUDED.login,"+
		"  password = EXCLUDED.password;", service.Name, service.Login, service.Password, service.UserChatID)
	return err
}

func GetService(serviceName string, user int64) (models.Service, error) {
	log.Print("Getting service...")
	var service models.Service
	query := `SELECT * FROM services WHERE name = $1 AND user_chat_id = $2`
	err := pool.QueryRow(context.Background(), query, serviceName, user).Scan(&service.ID, &service.Name, &service.Login, &service.Password, &service.UserChatID)
	return service, err
}

func DeleteService(serviceName string, user int64) error {
	query := `DELETE FROM services WHERE name = $1 AND user_chat_id = $2`
	result, err := pool.Exec(context.Background(), query, serviceName, user)
	if err != nil {
		return err
	}

	affectedRows := result.RowsAffected()

	if affectedRows == 0 {
		return errors.New("service not found")
	}

	return nil
}
