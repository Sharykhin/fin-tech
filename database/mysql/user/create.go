package user

import (
	"context"
	"fmt"
	"time"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/request"
)

// Create creates a new user in a database and return just created row as entity
func (s UserService) Create(ctx context.Context, ur request.UserCreateRequest) (*entity.User, error) {
	stm, err := s.db.PrepareContext(
		ctx,
		"INSERT INTO users (email, first_name, last_name, created_at) VALUES (?, ?, ?, ?)",
	)

	if err != nil {
		return nil, fmt.Errorf("could not create prepared statement: %v", err)
	}

	result, err := stm.Exec(ur.Email, ur.FirstName, ur.LastName, time.Now().UTC())
	if err != nil {
		return nil, fmt.Errorf("could not execute a prepared statement: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}

	u := entity.User{
		ID:        id,
		Email:     ur.Email,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		CreatedAt: entity.Time(time.Now().UTC()),
		DeletedAt: entity.NullTime{
			Valid: false,
			Time:  entity.Time{},
		},
	}

	return &u, nil
}
