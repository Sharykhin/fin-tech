package database

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database/mysql/broker"
	"github.com/Sharykhin/fin-tech/database/mysql/company"
	"github.com/Sharykhin/fin-tech/database/mysql/user"
)

var UserStorage contract.UserStorage = user.NewUserService()

// NewUserStorage return
func NewUserStorage() contract.UserStorage {
	return user.NewUserService()
}

// NewCompanyStorage returns a new instance of company storage
func NewCompanyStorage() *company.Storage {
	return company.NewStorage()
}

func NewBrokerStorage() *broker.Storage {
	return broker.NewStorage()
}
