package broker

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/http/errs"
)

// Get returns a row of broker
func (s Storage) Get(ctx context.Context, UserID int64) (*entity.Broker, error) {
	row := s.db.QueryRowContext(
		ctx,
		"SELECT u.id, u.email, b.position, b.role, b.created_at, b.deleted_at "+
			"FROM brokers AS b "+
			"INNER JOIN users AS u ON b.user_id = u.id "+
			"WHERE b.user_id = ?",
		UserID,
	)

	var ui entity.UserIdentity
	var b entity.Broker
	err := row.Scan(&ui.ID, &ui.Email, &b.Position, &b.Role, &b.CreatedAt, &b.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ResourceNotFound
		}

		return nil, fmt.Errorf("could not scan broker: %v", err)
	}

	b.User = ui

	return &b, nil
}
