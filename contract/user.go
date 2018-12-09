package contract

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

type (
	// UserStorage is an interface that specifies general CRUD operations
	// that can be performed around user entity
	UserStorage interface {
		Create(ctx context.Context, ur request.UserCreateRequest) (*entity.User, error)
	}
)
