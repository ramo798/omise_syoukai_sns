package db

import (
	"log"

	"github.com/ramo798/omise_syoukai_sns/model"
)

// PutterPost is 投稿の情報を受け取りデータベースに登録する。
func PutterPost(postinfo model.PostForRegister) {
	// exp INSERT INTO post_list(`users_id`, `comment`, `restaurant_id`, `price`, `assessment`) VALUES ('sas', 'sasa', 's1000', '1230', '50');

	db := initDb()
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO post_list(`users_id`, `comment`, `restaurant_id`, `price`, `assessment`) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	ins.Exec(postinfo.UsersID, postinfo.Comment, postinfo.RestaurantID, postinfo.Price, postinfo.Assessment)
}
