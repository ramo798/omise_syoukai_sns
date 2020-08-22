package db

import (
	"testing"

	"github.com/ramo798/omise_syoukai_sns/model"
)

func TestAbs(t *testing.T) {

	post := model.PostForRegister{
		UsersID:      "javatea",
		Comment:      "saiko",
		RestaurantID: "jssss",
		Price:        3400,
		Assessment:   70,
	}
	PutterPost(post)
	// GetterUserPost("ramo123")
}
