package views

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

			newEmployee.Id = employeeRecord.(dbtype.Node).Id
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
		positionRecord, ok := record.Get("positions")

		if !ok {
			continue
		}

		newPosition.Id = positionRecord.(dbtype.Node).Id
		props := positionRecord.(dbtype.Node).Props
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
		projectRecord, ok := record.Get("projects")

		if !ok {
			continue
		}

		newProject.Id = projectRecord.(dbtype.Node).Id
		props := projectRecord.(dbtype.Node).Props
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
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	params := mux.Vars(r)
	employeeId := params["employeeId"]
	query := fmt.Sprintf("MATCH (employee:Employee) WHERE ID(employee)=%v RETURN employee", employeeId)
	employee := models.Employee{}
	employee.Id = -1

	result, _ := session.Run(query, nil)
	record, err := result.Single()
	if err != nil {
		fmt.Errorf("not found: %v", employeeId)
		fmt.Errorf("message: %v", err)
	} else {
		employeeRecord, _ := record.Get("employee")
		employee.Id = employeeRecord.(dbtype.Node).Id
		props := employeeRecord.(dbtype.Node).Props
		name, ok := props["name"]
		if ok {
			employee.Name = name.(string)
		}
		surname, ok := props["surname"]
		if ok {
			employee.Surname = surname.(string)
		}
		phoneNumber, ok := props["phoneNumber"]
		if ok {
			employee.PhoneNumber = phoneNumber.(string)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	employeeId := params["employeeId"]
	query := fmt.Sprintf("MATCH (employee:Employee) WHERE ID(employee)=%v DELETE employee", employeeId)
	modificationStatus := models.ModificationStatus{}
	_, err := session.Run(query, nil)
	if err != nil {
		fmt.Errorf("not found: %v", employeeId)
		fmt.Errorf("message: %v", err)
		modificationStatus.Status = "not deleted"
		modificationStatus.Error = err.Error()
	} else {
		modificationStatus.Status = "ok"
		modificationStatus.Error = ""
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificationStatus)
}

func createPosition(w http.ResponseWriter, r *http.Request) {
}

func readPosition(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	params := mux.Vars(r)
	positionId := params["positionId"]
	query := fmt.Sprintf("MATCH (position:Position) WHERE ID(position)=%v RETURN position", positionId)
	position := models.Position{}
	position.Id = -1

	result, _ := session.Run(query, nil)
	record, err := result.Single()
	if err != nil {
		fmt.Errorf("not found: %v", positionId)
		fmt.Errorf("message: %v", err)
	} else {
		positionRecord, _ := record.Get("position")
		position.Id = positionRecord.(dbtype.Node).Id
		props := positionRecord.(dbtype.Node).Props
		name, ok := props["name"]
		if ok {
			position.Name = name.(string)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(position)
}

func updatePosition(w http.ResponseWriter, r *http.Request) {
}

func deletePosition(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	positionId := params["positionId"]
	query := fmt.Sprintf("MATCH (position:Position) WHERE ID(position)=%v DELETE position", positionId)
	modificationStatus := models.ModificationStatus{}
	result, _ := session.Run(query, nil)
	_, err := result.Single()
	if err != nil {
		fmt.Errorf("not found: %v", positionId)
		fmt.Errorf("message: %v", err)
		modificationStatus.Status = "not deleted"
		modificationStatus.Error = err.Error()
	} else {
		modificationStatus.Status = "ok"
		modificationStatus.Error = ""
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificationStatus)
}

func createProject(w http.ResponseWriter, r *http.Request) {

}

func readProject(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	params := mux.Vars(r)
	projectId := params["projectId"]
	query := fmt.Sprintf("MATCH (project:Project) WHERE ID(project)=%v RETURN project", projectId)
	project := models.Project{}
	project.Id = -1

	result, _ := session.Run(query, nil)
	record, err := result.Single()
	if err != nil {
		fmt.Errorf("not found: %v", projectId)
		fmt.Errorf("message: %v", err)
	} else {
		positionRecord, _ := record.Get("project")
		project.Id = positionRecord.(dbtype.Node).Id
		props := positionRecord.(dbtype.Node).Props
		name, ok := props["name"]
		if ok {
			project.Name = name.(string)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func updateProject(w http.ResponseWriter, r *http.Request) {
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	projectId := params["projectId"]
	query := fmt.Sprintf("MATCH (project:Project) WHERE ID(project)=%v DELETE project", projectId)
	modificationStatus := models.ModificationStatus{}
	_, err := session.Run(query, nil)
	if err != nil {
		fmt.Errorf("not found: %v", projectId)
		fmt.Errorf("message: %v", err)
		modificationStatus.Status = "not deleted"
		modificationStatus.Error = err.Error()
	} else {
		modificationStatus.Status = "ok"
		modificationStatus.Error = ""
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modificationStatus)
}
