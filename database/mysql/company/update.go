package company

import (
	"context"
	"fmt"
	"strings"

	"github.com/Sharykhin/fin-tech/http/errs"

	"github.com/Sharykhin/fin-tech/request"
)

// Update method run sql query update
func (s Storage) Update(ctx context.Context, ID int64, ucr request.UpdateCompanyRequest) error {
	query := "UPDATE companies SET %s WHERE id = ?"
	var replacements []interface{}
	var sets string
	if ucr.Symbol.Valid {
		sets += "symbol=?,"
		replacements = append(replacements, ucr.Symbol.Value)
	}

	if ucr.Name.Valid {
		sets += "name=?,"
		replacements = append(replacements, ucr.Name.Value)
	}

	if ucr.Exchange.Valid {
		sets += "exchange=?,"
		replacements = append(replacements, ucr.Exchange.Value)
	}

	if ucr.Website.Valid {
		sets += "website=?,"
		replacements = append(replacements, ucr.Website.Value)
	}

	if ucr.Description.Valid {
		sets += "description=?,"
		replacements = append(replacements, ucr.Description.Value)
	}

	if ucr.Tags.Valid {
		sets += "tags=?,"
		replacements = append(replacements, ucr.Tags.Value)
	}

	if len(replacements) == 0 {
		return errs.NothingToUpdate
	}

	replacements = append(replacements, ID)
	sets = strings.TrimRight(sets, ",")
	query = fmt.Sprintf(query, sets)
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("could not create a prepared statement: %v", err)
	}

	_, err = stmt.Exec(replacements...)
	if err != nil {
		return fmt.Errorf("could not execute prepared statement: %v", err)
	}

	return nil
}
