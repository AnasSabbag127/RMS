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
	sql := `INSERT INTO restraunts(name,address,created_by,star) VALUES($1,$2,$3,$4)`	
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
