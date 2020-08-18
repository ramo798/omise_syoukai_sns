package model

type Restaurant_info struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Logo_image_url string `json:"logo_image_url"`
	Station_name   string `json:"station_name"`
}
