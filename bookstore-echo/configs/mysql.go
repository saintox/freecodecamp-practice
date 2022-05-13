package configs

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	databaseConnection := "root:@tcp(localhost:3306)/bookstore_db?parseTime=true"
	db, err := sql.Open("mysql", databaseConnection)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
