package company

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
)

// Find finds a company by ID
func (cc CompanyController) Find(ctx context.Context, ID int64) (*entity.Company, error) {
	cc.logger.Info("[%s] Method %s is called with ID %d", "CompanyController", "Find", ID)
	return cc.storage.Find(ctx, ID)
}
