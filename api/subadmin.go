package api
import (
	"net/http"
	"goPractice/model"
	"encoding/json"
	"goPractice/database"
	"log"
	"github.com/gorilla/mux"
)
//subadmin Routes
func SubAdminRoutes(router *mux.Router) *mux.Router{
	router.Handle("/create-users",http.HandlerFunc(CreateUserHandler)).Methods("POST")
	router.Handle("/create-dish",http.HandlerFunc(CreateUserHandler)).Methods("POST")
	router.Handle("/create-restraunt",http.HandlerFunc(CreateRestrauntHandler)).Methods("POST")
	return router
}

func CreateSubAdminHandler(w http.ResponseWriter,r *http.Request) {
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