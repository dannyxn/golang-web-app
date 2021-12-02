package views

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"golang-web-app/models"
	"net/http"
)

var DbSession *neo4j.Session

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createEmployee(w, r)
	} else if r.Method == "GET" {
		readEmployee(w, r)
	} else if r.Method == "PUT" {
		updateEmployee(w, r)
	} else if r.Method == "DELETE" {
		deleteEmployee(w, r)
	} else {
	}
}

func PositionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createPosition(w, r)
	} else if r.Method == "GET" {
		readPosition(w, r)
	} else if r.Method == "PUT" {
		updatePosition(w, r)
	} else if r.Method == "DELETE" {
		deletePosition(w, r)
	} else {
	}
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createProject(w, r)
	} else if r.Method == "GET" {
		readProject(w, r)
	} else if r.Method == "PUT" {
		updateProject(w, r)
	} else if r.Method == "DELETE" {
		deleteProject(w, r)
	} else {
	}
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
}

func readEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	e := models.Employee{ID: params["employee_id"], Name: "Dawid", Surname: "Chara", PositionID: "1", PhoneNumber: "123132123"}
	json.NewEncoder(w).Encode(e)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
}

func createPosition(w http.ResponseWriter, r *http.Request) {
}

func readPosition(w http.ResponseWriter, r *http.Request) {
}

func updatePosition(w http.ResponseWriter, r *http.Request) {
}

func deletePosition(w http.ResponseWriter, r *http.Request) {
}

func createProject(w http.ResponseWriter, r *http.Request) {
}

func readProject(w http.ResponseWriter, r *http.Request) {
}

func updateProject(w http.ResponseWriter, r *http.Request) {
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
}
