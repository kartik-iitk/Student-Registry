package main

import (
	"log"
	"net/http"

	"Student-Registry/controllers"
	"Student-Registry/database"
	"Student-Registry/entity"

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
	router.HandleFunc("/get/{ID}", controllers.GetStudentByRollNo).Methods("GET")
	router.HandleFunc("/update/{ID}", controllers.UpdateStudentByRollNo).Methods("PUT")
	router.HandleFunc("/delete/{ID}", controllers.DeletStudentByRollNo).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "root",
			DB:         "student_db",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Student{})
}
