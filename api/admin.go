package api
import (
	"net/http"
	"goPractice/model"
	"goPractice/database"
	"encoding/json"
	"log"
)
//admin routes 


func CreateUserHandler(w http.ResponseWriter,r *http.Request) {
	var newUser model.InputUser
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newUser);err!=nil{
		http.Error(w,"Invalid Request Body",http.StatusBadRequest)
		return 
	}
	db,err := database.ConnectToDB()

	if err!=nil {
		http.Error(w,"failed to connect DB ",http.StatusInternalServerError)
		return 
	}
	sql := `INSERT INTO users(name,email,address,password,role_id) VALUES($1,$2,$3,$4,$5)`
	_,err = db.Exec(sql,newUser.Name,newUser.Email,newUser.Address,newUser.Password,newUser.RoleId)

	if err!=nil{
		log.Println("DATABASE ERROR: ",err)
		http.Error(w,"DATABASE ERROR: ",http.StatusInternalServerError)
		return 
	}
	response := map[string]interface{}{
		"message":"user created successfully",
		"user":newUser,
	}
	jsonResponse,err :=json.Marshal(response)
	if err!=nil {
		log.Println("json Marshaling error:  ",err)
		http.Error(w,"json marhsaling Error",http.StatusInternalServerError)
		return 
	} 
	w.Write(jsonResponse)

}	


func CreateDishesHandler(w http.ResponseWriter,r *http.Request){
	var newDish model.InputDishes
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newDish);err != nil{
		http.Error(w,"Invalid Request Body",http.StatusBadRequest)
		return 
	}
	db,err := database.ConnectToDB()

	if err!=nil {
		http.Error(w,"failed to connect DB ",http.StatusInternalServerError)
		return 
	}
	sql := `INSERT INTO dishes(restraunt_id,created_by,dish_name,price) VALUES($1,$2,$3.$4)`	
	_,err = db.Exec(sql,newDish.RestaurantId,newDish.CreatedBy,newDish.DishName,newDish.Price)

	if err!=nil{
		log.Println("DATABASE ERROR: ",err)
		http.Error(w,"DATABASE ERROR: ",http.StatusInternalServerError)
		return 
	}
	response := map[string]interface{}{
		"message":"Dish created successfully",
		"dish":newDish,
	}
	jsonResponse,err :=json.Marshal(response)
	if err!=nil {
		log.Println("json Marshaling error:  ",err)
		http.Error(w,"json marhsaling Error",http.StatusInternalServerError)
		return 
	} 
	w.Write(jsonResponse)

	
}



func CreateRestrauntHandler(w http.ResponseWriter,r *http.Request){
		
	var newRest model.InputRestraunt
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newRest);err != nil{
		http.Error(w,"Invalid Request Body",http.StatusBadRequest)
		return 
	}
	db,err := database.ConnectToDB()

	if err!=nil {
		http.Error(w,"failed to connect DB ",http.StatusInternalServerError)
		return 
	}
	sql := `INSERT INTO restraunts(name,address,created_by,stars) VALUES($1,$2,$3.$4)`	
	_,err = db.Exec(sql,newRest.Name,newRest.Address,newRest.CreatedBy,newRest.Stars)

	if err!=nil{
		log.Println("DATABASE ERROR: ",err)
		http.Error(w,"DATABASE ERROR: ",http.StatusInternalServerError)
		return 
	}
	response := map[string]interface{}{
		"message":"Restraunt created successfully",
		"Restraunt":newRest,
	}
	jsonResponse,err :=json.Marshal(response)
	if err!=nil {
		log.Println("json Marshaling error:  ",err)
		http.Error(w,"json marhsaling Error",http.StatusInternalServerError)
		return 
	} 
	w.Write(jsonResponse)
}



// // 
// func getAllUserHandler(w http.ResponseWriter,r *http.Request){
// 	sql = `SELECT id,name,address,role_id FROM users `

// }
// func getAllRestraunts(w http.ResponseWriter,r *http.Request){
// 	sql = `SELECT id,name,address,created_by,stars restraunt`
// }
// func getAllDishes(w http.ResponseWriter,r *http.Request){
// 	sql = `SELECT id,restraunt_id,created_by,dish_name,price FROM dishes`
// }



