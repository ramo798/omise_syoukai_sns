package db

import (
	"log"

	"github.com/ramo798/omise_syoukai_sns/model"
)

// GetterAllPost is すべてのポストを取得するための関数
func GetterAllPost() ([]model.Post, error) {
	db := initDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM post_list")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var posts []model.Post

	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.UsersID, &post.Comment, &post.RestaurantID, &post.Price, &post.Assessment)
		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	// log.Println("--------------", posts)

	return posts, nil
}

// GetterUserPost is 特定のユーザのポストリストをとってくる
func GetterUserPost() {
	db := initDb()
	defer db.Close()

	username := "ramo123"
	query := "SELECT * FROM post_list" + " WHERE " + " users_id = " + "'" + username + "'"
	log.Println(query)

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var posts []model.Post

	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.UsersID, &post.Comment, &post.RestaurantID, &post.Price, &post.Assessment)
		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	log.Println(posts)

	// log.Println("--------------", posts)

	// return posts, nil

}
