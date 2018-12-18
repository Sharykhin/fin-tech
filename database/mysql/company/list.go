package company

import (
	"context"
	"fmt"

	"github.com/Sharykhin/fin-tech/entity"
)

func (s Storage) List(ctx context.Context, limit, offset int64) ([]entity.Company, error) {
	rows, err := s.db.QueryContext(
		ctx,
		"SELECT id, symbol, name, exchange, website, description, tags FROM companies LIMIT ? OFFSET ?",
		limit,
		offset,
	)

	if err != nil {
		return nil, fmt.Errorf("could not execute sql query request: %v", err)
	}

	var companies []entity.Company
	for rows.Next() {
		var company entity.Company
		err := rows.Scan(&company.ID, &company.Symbol, &company.Name, &company.Exchange, &company.Website, &company.Description, &company.Tags)
		if err != nil {
			return nil, fmt.Errorf("could not scan a coompany row: %v", err)
		}
		companies = append(companies, company)
	}

	return companies, nil
}
