package company

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

func (c CompanyController) Create(ctx context.Context, ccr request.CreateCompanyRequest) (*entity.Company, error) {
	return c.storage.Create(ctx, ccr)
}
