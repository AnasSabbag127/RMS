package model

type Restaurant struct{
	Id int		`json:"id"`
	Name string	 `json:"RestaurantName"`
	Address string `json:"address"`
	CreatedBy int	`json:"createdBy"`
	Stars int `json:"stars"`
}

type InputRestraunt struct{
	Name string	 	`json:"name"`
	Address string `json:"RestrauntAddress"`
	CreatedBy int	`json:"createdBy"`
	Stars int  		`json:"stars"`
}
