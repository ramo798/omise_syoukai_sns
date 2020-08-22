package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ramo798/omise_syoukai_sns/db"
	"github.com/ramo798/omise_syoukai_sns/model"
)

func PostRegist(c *gin.Context) {
	var postinfo model.PostForRegister
	c.BindJSON(&postinfo)

	db.PutterPost(postinfo)

	c.Status(200)
}
