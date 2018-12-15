package company

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database"
)

type (
	Company struct {
		storage contract.CompanyStorage
	}
)

func NewCompany() *Company {
	return &Company{
		storage: database.NewCompanyStorage(),
	}
}
