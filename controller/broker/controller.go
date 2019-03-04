package broker

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/controller"
	"github.com/Sharykhin/fin-tech/database"
)

type (
	// Controller is a controller that would
	// handle business logic around broker entity
	Controller struct {
		controller.Controller
		storage contract.BrokerStorage
	}
)

// NewController is a function constructor
// that returns a new instance of BrokerController struct
func NewController() *Controller {
	return &Controller{
		Controller: controller.NewController(),
		storage: database.NewBrokerStorage(),
	}
}
