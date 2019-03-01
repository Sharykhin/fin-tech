package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	conn *sql.DB
)

func init() {

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&interpolateParams=false",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))

	if err != nil {
		panic(fmt.Errorf("could not open a connection to mysql: %v", err))
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	conn = db
}

// GetConnection returns mysql connection
func GetConnection() *sql.DB {
	return conn
}
