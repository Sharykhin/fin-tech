package broker

import "github.com/Sharykhin/fin-tech/controller/broker"

type HTTPHandler struct {
	controller *broker.BrokerController
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{
		controller: broker.NewBrokerController(),
	}
}