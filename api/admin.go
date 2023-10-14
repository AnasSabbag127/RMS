package api
import (
	"net/http"
	"github.com/gorilla/mux"
	// "strconv"
)
//admin routes 
func AdminRoutes(router *mux.Router) *mux.Router{
	router.Handle("/create-users",http.HandlerFunc(CreateUserHandler)).Methods("POST")
	router.Handle("/create-subadmin",http.HandlerFunc(CreateUserHandler)).Methods("POST")
	router.Handle("/create-restraunt",http.HandlerFunc(CreateRestrauntHandler)).Methods("POST")
	router.Handle("/create-dishes",http.HandlerFunc(CreateDishesHandler)).Methods("POST")
	return router
}

















