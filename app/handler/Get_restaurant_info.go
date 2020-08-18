package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ramo798/omise_syoukai_sns/db"
	"github.com/ramo798/omise_syoukai_sns/model"
)

func Get_restaurant_info(c *gin.Context) {

	var target_id model.Receiver_restaurant_info
	c.BindJSON(&target_id)
	// fmt.Println(target_id)
	a := db.Res_restaurant_info(target_id.Restaurant_id)
	// fmt.Println(a[0])

	tmp := a[0]

	// jsonBytes, err := json.Marshal(tmp)
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// res := string(jsonBytes)
	// fmt.Println(res)

	c.JSON(200, tmp)
}
