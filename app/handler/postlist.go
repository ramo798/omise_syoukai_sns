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
	log.Println("allpost-----", posts)
	// for _, j := range posts {
	// 	log.Printf("%#v", j)
	// }
	c.JSON(200, posts)
}

// GetUserPost is 特定のユーザのポスト一覧をjsonで返す。
func GetUserPost(c *gin.Context) {
	// http://~~~.com/post/xxx のxxxをストリングでとる処理
	username := c.Param("username")
	// log.Println("handler.postlist---username:", username)

	posts, err := db.GetterUserPost(username)
	if err != nil {
		log.Println("err")
	}
	// log.Println("handler.postlist---userpost:", posts)
	// for _, j := range posts {
	// 	log.Printf("%#v", j)
	// }
	c.JSON(200, posts)
}
