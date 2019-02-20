package broker

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/controller"
	"github.com/Sharykhin/fin-tech/database"
)

type (
	// BrokerController is a controller that would
	// handle business logic around broker entity
	BrokerController struct {
		controller.Controller
		storage contract.BrokerStorage
	}
)

// NewBrokerController is a function constructor
// that returns a new instance of BrokerController struct
func NewBrokerController() *BrokerController {
	ctrl := controller.NewController()
	return &BrokerController{
		Controller: ctrl,
		storage: database.NewBrokerStorage(),
	}
}
