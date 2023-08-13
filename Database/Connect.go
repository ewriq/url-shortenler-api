package Database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/urlshort")
	if err != nil {
		panic(err.Error())
	}
}
