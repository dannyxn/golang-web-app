package views

import (
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"golang-web-app/models"
	"log"
	"net/http"
)

var DbDriver *neo4j.Driver

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func ListEmployees(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	result, err := session.Run(`
    MATCH (employees:Employee) RETURN employees`, nil)

	if err != nil {
		log.Fatal(err)
	} else {
		var employees []models.Employee
		for result.Next() {
			newEmployee := models.Employee{}
			record := result.Record()
			employeeRecord, ok := record.Get("employees")

			if !ok {
				continue
			}

			props := employeeRecord.(dbtype.Node).Props
			name, ok := props["name"]
			if ok {
				newEmployee.Name = name.(string)
			}
			surname, ok := props["surname"]
			if ok {
				newEmployee.Surname = surname.(string)
			}
			phoneNumber, ok := props["phoneNumber"]
			if ok {
				newEmployee.PhoneNumber = phoneNumber.(string)
			}
			if (newEmployee != models.Employee{}) {
				employees = append(employees, newEmployee)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}

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

func ListPositions(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	result, _ := session.Run(`
    MATCH (positions:Position) RETURN positions`, nil)

	var positions []models.Position
	for result.Next() {
		newPosition := models.Position{}
		record := result.Record()
		employeeRecord, ok := record.Get("positions")

		if !ok {
			continue
		}

		props := employeeRecord.(dbtype.Node).Props
		name, ok := props["name"]
		if ok {
			newPosition.Name = name.(string)
		}
		if (newPosition != models.Position{}) {
			positions = append(positions, newPosition)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(positions)

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

func ListProjects(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	result, _ := session.Run(`
    MATCH (projects:Project) RETURN projects`, nil)

	var projects []models.Project
	for result.Next() {
		newProject := models.Project{}
		record := result.Record()
		employeeRecord, ok := record.Get("projects")

		if !ok {
			continue
		}

		props := employeeRecord.(dbtype.Node).Props
		name, ok := props["name"]
		if ok {
			newProject.Name = name.(string)
		}
		if (newProject != models.Project{}) {
			projects = append(projects, newProject)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
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
	//"CREATE (n:Employee {name: \"John\", surname: \"Doe\", phoneNumber: \"123123123\" })"

}

func readEmployee(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(e)
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
