package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramo798/omise_syoukai_sns/db"
	"github.com/ramo798/omise_syoukai_sns/model"
)

// GetRestaurantInfo is aaa
func GetRestaurantInfo(c *gin.Context) {

	var targetID model.ReceiverRestaurantInfo
	c.BindJSON(&targetID)
	// fmt.Println(target_id)
	a, err := db.ResRestaurantInfo(targetID.RestaurantID)
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
