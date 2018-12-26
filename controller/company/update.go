package company

import (
	"context"
	"fmt"
	"time"

	"github.com/Sharykhin/fin-tech/request"
)

func (cc CompanyController) Update(ctx context.Context, ID int64, ucr request.UpdateCompanyRequest) error {
	err := cc.storage.Update(ctx, ID, ucr)
	if err != nil {
		return err
	}

	go func() {
		time.After(5 * time.Second)
		fmt.Println("regenerate cache")
	}()

	return nil
}
