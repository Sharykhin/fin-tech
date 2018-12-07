package contract

import "context"

type (
	UserStorage interface {
		Create(ctx context.Context)
	}
)
