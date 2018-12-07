package user

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database/mysql/user"
)

type (
	UserController struct {
		storage contract.UserStorage
	}
)

// NewUserController returns a new instance of a controller
func NewUserController() *UserController {
	return &UserController{
		storage: user.NewCRUDService(),
	}
}
