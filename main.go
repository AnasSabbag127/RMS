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


// import (
// 	"fmt"
// 	// "github.com/gin-gonic/gin"
// 	// "net/http"
// )



// // type Task struct {
// // 	Id       int    `json:"id"`
// // 	TaskName string `json:"taskName"`
// // }

// // var tasks []Task

// // func main() {
// // 	fmt.Println("hii ")
// // 	// r := gin.Default()
// // 	// r.POST("/createTask", CreateTask)
// // 	// r.GET("/getAllTask", GetAllTasks)
// // 	// r.Run(":8000")
// // }

// // func CreateTask(c *gin.Context) {
// // 	var newTask Task

// // 	if err := c.BindJSON(&newTask); err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{"err": err})
// // 		return
// // 	}

// // 	newTask.Id = len(tasks) + 1
// // 	tasks = append(tasks, newTask)
// // 	c.JSON(http.StatusCreated, gin.H{"task": newTask})
// // }

// // func GetAllTasks(c *gin.Context) {
// // 	c.JSON(http.StatusOK, tasks)
// // }
