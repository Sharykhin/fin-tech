package company

import (
	"database/sql"

	"github.com/Sharykhin/fin-tech/database/mysql"
)

type Storage struct {
	db *sql.DB
}

// NewStorage returns a new instance of company storage service
func NewStorage() *Storage {
	return &Storage{
		db: mysql.GetConnection(),
	}
}
