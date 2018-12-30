package company

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Sharykhin/fin-tech/http/errs"

	"github.com/Sharykhin/fin-tech/entity"
)

func (s Storage) Find(ctx context.Context, ID int64) (*entity.Company, error) {
	row := s.db.QueryRowContext(
		ctx,
		"SELECT id, symbol, name, exchange, website, description, tags FROM companies WHERE id = ?",
		ID,
	)

	var company entity.Company
	err := row.Scan(
		&company.ID,
		&company.Symbol,
		&company.Name,
		&company.Exchange,
		&company.Website,
		&company.Description,
		&company.Tags,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ResourceNotFound
		}
		return nil, fmt.Errorf("could not scan company: %v", err)
	}

	return &company, nil
}
