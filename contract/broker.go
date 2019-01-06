package contract

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

const (
	ROLE_DEFENDER = "defender"
)

type (
	BrokerStorage interface {
		Get(ctx context.Context, UserID int64) (*entity.Broker, error)
		Update(ctx context.Context, ID int64, ubr request.UpdateBrokerRequest) error
	}
)
