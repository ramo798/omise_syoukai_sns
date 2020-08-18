package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramo798/omise_syoukai_sns/db"
	"github.com/ramo798/omise_syoukai_sns/model"
)

func Get_restaurant_info(c *gin.Context) {

	var target_id model.Receiver_restaurant_info
	c.BindJSON(&target_id)
	// fmt.Println(target_id)
	a, err := db.Res_restaurant_info(target_id.Restaurant_id)
	if err != nil {
		log.Println("aaa")
		c.String(510, "tmp")
		return
	}
	// fmt.Println(a[0])

	if len(a) != 0 {
		tmp := a[0]
		c.JSON(200, tmp)
		return
	}

	c.Status(500)

}
