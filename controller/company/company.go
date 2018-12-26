package company

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database"
	"github.com/Sharykhin/fin-tech/service/logger"
)

type (
	CompanyController struct {
		storage contract.CompanyStorage
		logger  contract.Logger
	}
)

func NewCompanyController() *CompanyController {
	return &CompanyController{
		storage: database.NewCompanyStorage(),
		logger:  logger.Logger,
	}
}
