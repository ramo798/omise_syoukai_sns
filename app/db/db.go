package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Hello is test
func Hello() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql_host:3306)/mydatabase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
