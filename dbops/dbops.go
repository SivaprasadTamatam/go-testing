package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetupDB() {
	dsn := "user:password@tcp(localhost:3306)/mydb"

	var err error
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
}

func TeardownDB() {
	if db != nil {
		db.Close()
	}
}
