package controller

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/service/logger"
)

// Controller is a general struct that would be embed into specific controllers
type Controller struct {
	logger  contract.Logger
}

// NewController is a constructor function that returns a new instance of a controller struct
// with injected dependencies
func NewController() Controller {
	return Controller{
		logger: logger.Logger,
	}
}
