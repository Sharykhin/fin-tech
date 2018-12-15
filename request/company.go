package request

import "github.com/Sharykhin/fin-tech/entity"

type (
	CreateCompanyRequest struct {
		Symbol      string      `json:"symbol"`
		Name        string      `json:"name"`
		Exchange    string      `json:"exchange"`
		Website     string      `json:"website"`
		Description string      `json:"description"`
		Tags        entity.Tags `json:"tags"`
	}
)
