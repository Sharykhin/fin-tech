package user

import (
	"database/sql"

	"github.com/Sharykhin/fin-tech/database/mysql"
)

type (
	// UserService provides basic crud operation around User struct
	UserService struct {
		db *sql.DB
	}
)

// NewUserService returns a new instance of UserService struct
func NewUserService() *UserService {
	us := UserService{
		db: mysql.GetConnection(),
	}

	return &us
}
