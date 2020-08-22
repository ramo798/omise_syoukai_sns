package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramo798/omise_syoukai_sns/db"
)

// GetAllPost is すべてのポストをjsonで返す。
func GetAllPost(c *gin.Context) {
	posts, err := db.GetterAllPost()
	if err != nil {
		log.Println("err")
	}
	for _, j := range posts {
		log.Printf("%#v", j)
	}
	c.JSON(200, posts)
}

// GetUserPost is 特定のユーザのポスト一覧をjsonで返す。
func GetUserPost() {
	// func GetUserPost(c *gin.Context) {
	// http://~~~.com/post/xxx のxxxをストリングでとる処理
	// username := c.Param("username")
	username := "ramo123"
	log.Println("username")

	posts, err := db.GetterUserPost(username)
	log.Println("posts")
	if err != nil {
		log.Println("err")
	}
	for _, j := range posts {
		log.Printf("%#v", j)
	}
	// c.JSON(200, posts)
}
