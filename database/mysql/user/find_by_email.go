package user

import (
	"context"
	"database/sql"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/http/errs"
)

func (us UserService) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	row := us.db.QueryRowContext(
		ctx,
		"SELECT id, email, first_name, last_name, created_at, deleted_at FROM users WHERE email=?",
		email,
	)
	err := row.Scan(user.ID, user.Email, user.FirstName, user.LastName, user.CreatedAt, user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.UserWasNotFound
		}

		return nil, err
	}

	return &user, nil
}
