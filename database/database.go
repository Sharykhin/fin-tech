package database

import (
	"context"
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database/mysql/broker"
	"github.com/Sharykhin/fin-tech/database/mysql/company"
	"github.com/Sharykhin/fin-tech/database/mysql/user"
)

type (
	UserRetriever interface {
		GetList(ctx context.Context, offset, limit int64) ([]string, error)
	}
)

var UserStorage contract.UserStorage = user.NewUserService()

// TODO: interfaces of storages should be place here, it would be more semantic
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
