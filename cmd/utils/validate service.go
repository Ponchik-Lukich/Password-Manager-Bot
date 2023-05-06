package utils

import (
	"errors"
	"log"
	"password-manager/cmd/models"
	"strings"
)

func ValidateService(credentials string, user int64) (models.Service, error) {
	var name, login, password string
	name = credentials[:strings.Index(credentials, ":")]
	if name == "" {
		log.Print("name is empty")
		return models.Service{}, errors.New("name is empty")
	}
	login = credentials[strings.Index(credentials, ":")+1 : strings.LastIndex(credentials, ":")]
	if login == "" {
		log.Print("login is empty")
		return models.Service{}, errors.New("login is empty")
	}
	password = credentials[strings.LastIndex(credentials, ":")+1:]
	if password == "" {
		log.Print("password is empty")
		return models.Service{}, errors.New("password is empty")
	}
	return models.Service{Name: name, Login: login, Password: password, UserChatID: user}, nil
}
