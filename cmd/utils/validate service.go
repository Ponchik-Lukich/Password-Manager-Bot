package utils

import (
	"errors"
	"password-manager/cmd/models"
	"regexp"
)

func ValidateService(credentials string, user int64) (models.Service, error) {
	var name, login, password string
	re := regexp.MustCompile(`^(\w+) (\w+) (\w+)$`)
	if !re.MatchString(credentials) {
		return models.Service{}, errors.New("invalid service credentials")
	} else {
		matches := re.FindStringSubmatch(credentials)
		name = matches[1]
		login = matches[2]
		password = matches[3]
	}
	if len(name) > 255 || len(login) > 255 || len(password) > 255 || len(name) == 0 || len(login) == 0 || len(password) == 0 {
		return models.Service{}, errors.New("service credentials are too long")
	} else {
		var err error
		login, err = Encrypt(login)
		if err != nil {
			return models.Service{}, err
		}
		password, err = Encrypt(password)
		if err != nil {
			return models.Service{}, err
		}
	}
	return models.Service{Name: name, Login: login, Password: password, UserChatID: user}, nil
}
