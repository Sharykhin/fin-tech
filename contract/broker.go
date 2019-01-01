package contract

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
)

const (
	ROLE_DEFENDER = "defender"
)

type (
	BrokerStorage interface {
		Get(ctx context.Context, UserID int64) (*entity.Broker, error)
	}
)
