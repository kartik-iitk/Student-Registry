package main

import (
	"log"
	"net/http"

	"github.com/kartik-iitk/Student-Registry/controllers"
	"github.com/kartik-iitk/Student-Registry/database"
	"github.com/kartik-iitk/Student-Registry/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllStudents).Methods("GET")
	router.HandleFunc("/get/{rollNo}", controllers.GetStudentByRollNo).Methods("GET")
	router.HandleFunc("/update/{rollNo}", controllers.UpdateStudentByRollNo).Methods("PUT")
	router.HandleFunc("/delete/{rollNo}", controllers.DeletStudentByRollNo).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "<username>",      // Enter your MySQL username.
			Password:   "<password>",      // Enter your MySQL password.
			DB:         "<database_name>", // Enter the database name created.
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Student{})
}
