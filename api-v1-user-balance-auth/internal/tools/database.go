package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username string
}

type BalanceDetails struct{
	Balance int64
	Username string
}

type DatabaseInterface interface{
	GetUserLoginDetails(Username string) *LoginDetails
	GetUserBalance(username string) *BalanceDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}