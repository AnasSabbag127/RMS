package database
import (
	"fmt"
	"log"
	"database/sql"
	"os"
	"github.com/joho/godotenv"
	
)
func ConnectToDB()(*sql.DB, error){
	err := godotenv.Load()
	if err != nil {
		return nil,err
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db,err := sql.Open("postgres",connString)
	if err!=nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("connect To DATABASE SUCCESSFULL ")
	return db,nil
}