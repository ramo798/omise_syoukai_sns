package db

import (
	"database/sql"
	"log"

	"github.com/ramo798/omise_syoukai_sns/model"
	// aaa
	_ "github.com/go-sql-driver/mysql"
)

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(mysql_host:3306)/mydatabase")
	if err != nil {
		panic(err.Error())
	}

	// log.Println("-----", reflect.TypeOf(db))
	return db
}

// ResRestaurantInfo is レストランIDを受け取ってRestaurant_infoを返す
func ResRestaurantInfo(restaurantID string) ([]model.RestaurantInfo, error) {
	db, err := sql.Open("mysql", "root:root@tcp(mysql_host:3306)/mydatabase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM restaurant_info WHERE ID = ?", restaurantID)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var res []model.RestaurantInfo

	for rows.Next() {
		var info model.RestaurantInfo
		err := rows.Scan(&info.ID, &info.Name, &info.Address, &info.LogoImageURL, &info.StationName)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(info.ID, info.Name, info.Address, info.Logo_image_url, info.Station_name)
		// fmt.Println(reflect.TypeOf(info))
		res = append(res, info)
	}

	if len(res) == 0 {
		log.Println("検索結果はありません。")
		return nil, err
	}

	return res, nil

}
