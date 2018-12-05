package user

import "database/sql"

type (
	// UserService provides basic crud operation around User struct
	UserService struct {
		db *sql.DB
	}
)
