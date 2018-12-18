package company

import (
	"context"
	"fmt"
)

// Count returns a total number of all rows
func (s Storage) Count(ctx context.Context) (int64, error) {
	row := s.db.QueryRowContext(ctx, "SELECT COUNT(id) FROM companies")
	var total int64
	err := row.Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("could not scan count query: %v", err)
	}

	return total, nil
}
