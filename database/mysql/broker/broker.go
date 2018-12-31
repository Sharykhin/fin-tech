package broker

import (
	"database/sql"

	"github.com/Sharykhin/fin-tech/database/mysql"
)

type Storage struct {
	db *sql.DB
}

// NewStorage returns a new instance of Storage struct
func NewStorage() *Storage {
	return &Storage{
		db: mysql.GetConnection(),
	}
}
