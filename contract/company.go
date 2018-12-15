package contract

import (
	"context"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

type (
	CompanyStorage interface {
		Create(context.Context, request.CreateCompanyRequest) (*entity.Company, error)
	}
)
