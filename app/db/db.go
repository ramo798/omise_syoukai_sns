package db

import (
	"database/sql"
	"fmt"

	"github.com/ramo798/omise_syoukai_sns/model"

	_ "github.com/go-sql-driver/mysql"
)

// Hello is test
func Hello() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql_host:3306)/mydatabase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM restaurant_info")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var info model.Restaurant_info
		err := rows.Scan(&info.ID, &info.Name, &info.Address, &info.Logo_image_url, &info.Station_name)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(info.ID, info.Name, info.Address, info.Logo_image_url, info.Station_name)

	}
}
