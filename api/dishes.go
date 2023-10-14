package api
import (
	"net/http"
	"goPractice/model"
	"goPractice/database"
	"encoding/json"
	"log"
	// "github.com/gorilla/mux"
	// "strconv"
)

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
	sql := `INSERT INTO dishes(restraunt_id,created_by,name,price) VALUES($1,$2,$3,$4)`	
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
