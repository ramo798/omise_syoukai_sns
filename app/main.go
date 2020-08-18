package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ramo798/omise_syoukai_sns/db"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello, API")
	})

	router.Run(":8080")
	db.Test()
}
