package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ramo798/omise_syoukai_sns/handler"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello, API")
	})
	router.POST("/getrestaurantinfo", handler.GetRestaurantInfo)

	router.GET("/postlist/all", handler.GetAllPost)

	router.GET("/postlist/user/:username", handler.GetUserPost)

	router.POST("/post/register", handler.PostRegist)

	router.Run(":8080")

}
