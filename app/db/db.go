package db

import (
	"database/sql"

	"github.com/ramo798/omise_syoukai_sns/model"

	_ "github.com/go-sql-driver/mysql"
)

// Res_restaurant_info is レストランIDを受け取ってRestaurant_infoを返す
func Res_restaurant_info(restaurant_id string) []model.Restaurant_info {
	db, err := sql.Open("mysql", "root:root@tcp(mysql_host:3306)/mydatabase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM restaurant_info WHERE ID = ?", restaurant_id)

	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var res []model.Restaurant_info

	for rows.Next() {
		var info model.Restaurant_info
		err := rows.Scan(&info.ID, &info.Name, &info.Address, &info.Logo_image_url, &info.Station_name)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(info.ID, info.Name, info.Address, info.Logo_image_url, info.Station_name)
		// fmt.Println(reflect.TypeOf(info))
		res = append(res, info)
	}

	return res

}
