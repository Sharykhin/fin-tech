package company

import (
	"database/sql"

	"github.com/Sharykhin/fin-tech/database/mysql"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {
	return &Storage{
		db: mysql.GetConnection(),
	}
}
