package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ramo798/omise_syoukai_sns/handler"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello, API")
	})
	router.POST("/getrestaurantinfo", handler.Get_restaurant_info)

	router.GET("/postlist/all", handler.Get_all_post)

	router.Run(":8080")

}
