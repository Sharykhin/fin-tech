package contract

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

type (
	CompanyStorage interface {
		Create(ctx context.Context, ccr request.CreateCompanyRequest) (*entity.Company, error)
		List(ctx context.Context, limit, offset int64) ([]entity.Company, error)
		Count(ctx context.Context) (int64, error)
	}
)
