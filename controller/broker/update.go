package broker

import (
	"context"

	"github.com/Sharykhin/fin-tech/request"
)

func (bc Controller) Update(ctx context.Context, ID int64, ubr request.UpdateBrokerRequest) error {
	return bc.storage.Update(ctx, ID, ubr)
}
