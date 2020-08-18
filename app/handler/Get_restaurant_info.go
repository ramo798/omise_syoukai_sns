package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ramo798/omise_syoukai_sns/db"
	"github.com/ramo798/omise_syoukai_sns/model"
)

func Get_restaurant_info(c *gin.Context) {

	var target_id model.Receiver_restaurant_info
	c.BindJSON(&target_id)
	a := db.Res_restaurant_info(target_id.Restaurant_id)
	fmt.Println(a)

	// res := "eee"
	// c.String(200, res)
}
