package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ramo798/omise_syoukai_sns/db"
)

func main() {
	db.Hello()
}
