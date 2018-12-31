package company

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database"
	"github.com/Sharykhin/fin-tech/service/logger"
)

type (
	// CompanyController is a general controller around Company entity
	CompanyController struct {
		storage contract.CompanyStorage
		logger  contract.Logger
	}
)

// NewCompanyController returns a new instance of CompanyController
func NewCompanyController() *CompanyController {
	return &CompanyController{
		storage: database.NewCompanyStorage(),
		logger:  logger.Logger,
	}
}
