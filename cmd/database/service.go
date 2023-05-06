package database

import (
	"context"
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

func GetService(serviceName string) (models.Service, error) {
	var service models.Service
	query := `SELECT * FROM services WHERE name = $1`
	err := pool.QueryRow(context.Background(), query, serviceName).Scan(&service.ID, &service.Name, &service.Login, &service.Password, &service.UserChatID)
	return service, err
}

func DeleteService(serviceName string, user int64) error {
	query := `DELETE FROM services WHERE name = $1 AND user_chat_id = $2`
	_, err := pool.Exec(context.Background(), query, serviceName, user)
	return err
}
