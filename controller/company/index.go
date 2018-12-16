package company

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
)

func (c CompanyController) Index(ctx context.Context, limit, offset int64) ([]entity.Company, int64, error) {
	return make([]entity.Company, 0), 10, nil
}
