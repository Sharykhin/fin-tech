package database

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database/mysql/user"
)

var UserStorage contract.UserStorage = user.NewUserService()

// NewUserStorage return
func NewUserStorage() contract.UserStorage {
	return user.NewUserService()
}
