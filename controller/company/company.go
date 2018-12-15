package company

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database"
)

type (
	CompanyController struct {
		storage contract.CompanyStorage
	}
)

func NewCompanyController() *CompanyController {
	return &CompanyController{
		storage: database.NewCompanyStorage(),
	}
}
