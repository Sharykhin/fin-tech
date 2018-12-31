package broker

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/database"
	"github.com/Sharykhin/fin-tech/service/logger"
)

type (
	BrokerController struct {
		storage contract.BrokerStorage
		logger  contract.Logger
	}
)

func NewBrokerController() *BrokerController {
	return &BrokerController{
		storage: database.NewBrokerStorage(),
		logger:  logger.Logger,
	}
}
