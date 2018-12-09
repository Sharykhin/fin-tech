package user

import (
	"context"

	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database/mysql/user"
	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

// NewUserController returns a new instance of a controller
func NewUserController() *UserController {
	return &UserController{
		storage: user.NewCRUDService(),
	}
}

type (
	// UserController is a controller that provides necessary methods around User entity
	UserController struct {
		storage contract.UserStorage
	}
)

func (uc UserController) Create(ctx context.Context, ur request.UserCreateRequest) (*entity.User, error) {
	u, err := uc.storage.Create(ctx, ur)
	// TODO: Put logging here
	if err != nil {
		return u, err
	}

	return u, nil
}
