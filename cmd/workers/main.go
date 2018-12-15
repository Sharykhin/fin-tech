package main

import (
	"context"
	"fmt"

	"github.com/Sharykhin/fin-tech/entity"

	"github.com/Sharykhin/fin-tech/controller/company"
	"github.com/Sharykhin/fin-tech/request"
)

func main() {
	ccr := request.CreateCompanyRequest{
		Symbol:      "AAPL",
		Name:        "Apple Inc.",
		Exchange:    "Nasdaq Global Select",
		Website:     "http://www.apple.com",
		Description: "Apple Inc is designs, manufactures and markets mobile communication and media devices and personal computers, and sells a variety of related software, services, accessories, networking solutions and third-party digital content and applications.",
		Tags:        entity.Tags{"Technology", "Consumer Electronics", "Computer Hardware"},
	}

	ctrl := company.NewCompanyController()

	c, err := ctrl.Create(context.Background(), ccr)
	fmt.Println(c, err)
}
