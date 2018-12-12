package user

import (
	"context"

	"github.com/Sharykhin/fin-tech/database"

	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

// NewController returns a new instance of a controller
func NewController() *Controller {
	return &Controller{
		storage: database.NewUserStorage(),
	}
}

type (
	// UserController is a controller that provides necessary methods around User entity
	Controller struct {
		storage contract.UserStorage
	}
)

func (c Controller) Create(ctx context.Context, ur request.UserCreateRequest) (*entity.User, error) {
	u, err := c.storage.Create(ctx, ur)
	// TODO: Put logging here
	if err != nil {
		return u, err
	}

	return u, nil
}
