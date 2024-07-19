package tools

import (
	log "github.com/sirupsen/logrus"
)

// DB details
type LoginDetails struct {
	Username   string
	Auth_token string
}

type MoneyDetails struct {
	Username string
	Money    int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserMoney(username string) *MoneyDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	err := database.SetupDatabase()

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
