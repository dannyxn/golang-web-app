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

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

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
		respondWithJSON(w, 200, employees)
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
	respondWithJSON(w, 200, positions)
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
	respondWithJSON(w, 200, projects)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
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
		respondWithError(w, 404, fmt.Sprintf("Employee with id=%v not found", employeeId))
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
	respondWithJSON(w, 200, employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	employeeId := params["employeeId"]
	query := fmt.Sprintf("MATCH (employee:Employee) WHERE ID(employee)=%v DELETE employee", employeeId)
	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, 400, err.Error())
	} else {
		respondWithJSON(w, 200, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func CreatePosition(w http.ResponseWriter, r *http.Request) {
}

func GetPosition(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	params := mux.Vars(r)
	positionId := params["positionId"]
	query := fmt.Sprintf("MATCH (position:Position) WHERE ID(position)=%v RETURN position", positionId)
	position := models.Position{}
	position.Id = -1

	result, _ := session.Run(query, nil)
	record, err := result.Single()
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Position with id=%v not found", positionId))
	} else {
		positionRecord, _ := record.Get("position")
		position.Id = positionRecord.(dbtype.Node).Id
		props := positionRecord.(dbtype.Node).Props
		name, ok := props["name"]
		if ok {
			position.Name = name.(string)
		}
	}
	respondWithJSON(w, 200, position)
}

func UpdatePosition(w http.ResponseWriter, r *http.Request) {
}

func DeletePosition(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	positionId := params["positionId"]
	query := fmt.Sprintf("MATCH (position:Position) WHERE ID(position)=%v DELETE position", positionId)
	result, _ := session.Run(query, nil)
	_, err := result.Single()
	if err != nil {
		fmt.Errorf("not found: %v", positionId)
		fmt.Errorf("message: %v", err)
		respondWithError(w, 400, err.Error())
	} else {
		respondWithJSON(w, 200, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {

}

func GetProject(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	params := mux.Vars(r)
	projectId := params["projectId"]
	query := fmt.Sprintf("MATCH (project:Project) WHERE ID(project)=%v RETURN project", projectId)
	project := models.Project{}
	project.Id = -1

	result, _ := session.Run(query, nil)
	record, err := result.Single()
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Project with id=%v not found", projectId))

	} else {
		positionRecord, _ := record.Get("project")
		project.Id = positionRecord.(dbtype.Node).Id
		props := positionRecord.(dbtype.Node).Props
		name, ok := props["name"]
		if ok {
			project.Name = name.(string)
		}
	}
	respondWithJSON(w, 200, project)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	projectId := params["projectId"]
	query := fmt.Sprintf("MATCH (project:Project) WHERE ID(project)=%v DELETE project", projectId)
	_, err := session.Run(query, nil)
	if err != nil {
		fmt.Errorf("not found: %v", projectId)
		fmt.Errorf("message: %v", err)
		respondWithError(w, 400, err.Error())
	} else {
		respondWithJSON(w, 200, models.ModificationStatus{Status: "ok", Error: ""})
	}
}
