package main
import (
	"fmt"
	"goPractice/database"
	"goPractice/api"
	_ "github.com/lib/pq"
	"net/http"
)
func handler(_ http.ResponseWriter,r *http.Request){

}
func main(){
	fmt.Println("jfjkf")
	_,err := database.ConnectToDB()
	if err!=nil {
		fmt.Println("database not connected: Error: ",err)
		return 
	}
	// http.HandleFunc("/",handler)
	http.HandleFunc("/login",handler)
	//admin routes
	http.HandleFunc("/create-user",api.CreateUserHandler)
	http.HandleFunc("/create-restraunt",api.CreateRestrauntHandler)
	http.HandleFunc("/create-dishes",api.CreateDishesHandler)


	http.ListenAndServe("localhost:8000",nil)
	
}

