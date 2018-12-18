package company

import (
	"context"
	"fmt"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

// Create creates a new company row in a database
func (s Storage) Create(ctx context.Context, ccr request.CreateCompanyRequest) (*entity.Company, error) {
	stm, err := s.db.PrepareContext(
		ctx,
		"INSERT INTO companies (symbol, name, exchange, website, description, tags) VALUES (?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return nil, fmt.Errorf("could not create a prepared statement: %v", err)
	}

	result, err := stm.Exec(ccr.Symbol, ccr.Name, ccr.Exchange, ccr.Website, ccr.Description, ccr.Tags)
	if err != nil {
		return nil, fmt.Errorf("could not execute a prepared statement: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("coul not get last instert id: %v", err)
	}

	c := entity.Company{
		ID:          id,
		Symbol:      ccr.Symbol,
		Name:        ccr.Name,
		Exchange:    ccr.Exchange,
		Website:     ccr.Website,
		Description: ccr.Description,
		Tags:        ccr.Tags,
	}

	return &c, nil
}
