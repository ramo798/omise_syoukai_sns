package model

// Post is 投稿用の構造体
type Post struct {
	ID           int    `json:"post_id"`
	UsersID      string `json:"users_id"`
	Comment      string `json:"comment"`
	RestaurantID string `json:"sestaurant_id"`
	Price        int    `json:"price"`
	Assessment   int    `json:"assessment"`
}
