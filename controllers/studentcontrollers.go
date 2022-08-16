package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kartik-iitk/Student-Registry/database"
	"github.com/kartik-iitk/Student-Registry/entity"

	"github.com/gorilla/mux"
)

// GetAllStudents get all student data
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var persons []entity.Student
	database.Connector.Find(&persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

// GetStudentByRollNo returns student with specific RollNo
func GetStudentByRollNo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["rollNo"]

	var student entity.Student
	rollNo, _ := strconv.ParseInt(key, 10, 64)
	//database.Connector.QueryExpr(fmt.Sprintf("SELECT * FROM students WHERE roll_no = ?", key));
	database.Connector.Where("roll_no = ?", rollNo).Find(&student)

	//database.Connector.First(&student, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// CreateStudent creates students
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var student entity.Student
	json.Unmarshal(requestBody, &student)

	database.Connector.Create(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// UpdateStudentByRollNo updates student with respective RollNo
func UpdateStudentByRollNo(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var student entity.Student
	json.Unmarshal(requestBody, &student)

	database.Connector.Where("roll_no = ?", student.RollNo).Save(&student)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

// DeletStudentByRollNo delete's student with specific RollNo
func DeletStudentByRollNo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["rollNo"]

	var student entity.Student
	rollNo, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("roll_no = ?", rollNo).Delete(&student)
	w.WriteHeader(http.StatusNoContent)
}
