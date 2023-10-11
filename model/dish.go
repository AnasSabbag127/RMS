package model

type Dishes struct {
	Id int `json:"id"`
	RestaurantId int `json:"restrauntId"`
	CreatedBy int 	`json:"createdBy"`
	DishName string	`json:"dishName"`
	Price int `json:"price"`
}

type InputDishes struct{	
	RestaurantId int `json:"restrauntId"`
	CreatedBy int 	`json:"createdBy"`
	DishName string	`json:"dishName"`
	Price int 		`json:"price"`
}