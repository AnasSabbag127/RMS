package routes
import (
	"github.com/gorilla/mux"
	"goPractice/api"
	"goPractice/middlewares"
	"net/http"
)
func CreateRoutes() http.Handler{
	r := mux.NewRouter()
	r.Handle("/",api.AdminRoutes(r))
	r.Handle("/",api.SubAdminRoutes(r))
	r.Handle("/",api.UserRoutes(r))
	return middlewares.EnableCORS(r)
}
