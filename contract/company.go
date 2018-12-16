package contract

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

type (
	CompanyStorage interface {
		Create(ctx context.Context, ccr request.CreateCompanyRequest) (*entity.Company, error)
		//Index(ctx context.Context, limit, offset int64) ([]entity.Company, int64, error)
	}
)
