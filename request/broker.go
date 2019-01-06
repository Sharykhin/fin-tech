package request

import "github.com/Sharykhin/fin-tech/entity"

type (
	UpdateBrokerRequest struct {
		Position entity.NullInt    `json:"position"`
		Role     entity.NullString `json:"role"`
	}
)
