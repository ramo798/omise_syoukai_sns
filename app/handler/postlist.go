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
