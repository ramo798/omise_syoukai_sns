package model

//RestaurantInfo is aa
type RestaurantInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	LogoImageURL string `json:"logo_image_url"`
	StationName  string `json:"station_name"`
}
