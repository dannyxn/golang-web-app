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

func ListWorksIn(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	result, _ := session.Run(`
    MATCH p=()-[r:WorksIn]->() RETURN p`, nil)

	var worksInRelationships []models.WorksIn
	for result.Next() {
		newWorksInRelationship := models.WorksIn{}
		record := result.Record()
		result, ok := record.Get("p")
		path := result.(dbtype.Path)
		if ok {
			newWorksInRelationship.EmployeeId = path.Relationships[0].StartId
			newWorksInRelationship.ProjectId = path.Relationships[0].EndId
			worksInRelationships = append(worksInRelationships, newWorksInRelationship)
		}

	}
	respondWithJSON(w, 200, worksInRelationships)
}

func ListWorksAs(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	result, _ := session.Run(`
    MATCH p=()-[r:WorksAs]->() RETURN p`, nil)

	var worksAsRelationships []models.WorksAs
	for result.Next() {
		newWorksAsRelationship := models.WorksAs{}
		record := result.Record()
		result, ok := record.Get("p")
		path := result.(dbtype.Path)
		if ok {
			newWorksAsRelationship.EmployeeId = path.Relationships[0].StartId
			newWorksAsRelationship.PositionId = path.Relationships[0].EndId
			worksAsRelationships = append(worksAsRelationships, newWorksAsRelationship)
		}

	}
	respondWithJSON(w, 200, worksAsRelationships)
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

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	employeeId := params["employeeId"]
	query := fmt.Sprintf("MATCH (employee:Employee) WHERE ID(employee)=%v DETACH DELETE employee", employeeId)
	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, 400, err.Error())
	} else {
		respondWithJSON(w, 200, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeId := params["employeeId"]
	queryMatch := fmt.Sprintf("MATCH (employee:Employee) WHERE ID(employee)=%v", employeeId)
	querySet := " SET"

	var employee models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if employee.Name != "" {
		querySet += fmt.Sprintf(" employee.name = \"%v\"", employee.Name)
	}

	if employee.Surname != "" {
		querySet += fmt.Sprintf(" employee.surname = \"%v\"", employee.Surname)
	}

	if employee.PhoneNumber != "" {
		querySet += fmt.Sprintf(" employee.phoneNumber = \"%v\"", employee.PhoneNumber)
	}
	query := queryMatch + querySet + " RETURN employee"

	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	query := fmt.Sprintf("CREATE (employee:Employee {name: \"%v\", surname: \"%v\", phoneNumber: \"%v\"})",
		employee.Name, employee.Surname, employee.PhoneNumber)

	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
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
	params := mux.Vars(r)
	positionId := params["positionId"]
	queryMatch := fmt.Sprintf("MATCH (position:Position) WHERE ID(position)=%v", positionId)
	querySet := " SET"

	var position models.Position
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&position); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if position.Name != "" {
		querySet += fmt.Sprintf(" position.name = \"%v\"", position.Name)
	}

	query := queryMatch + querySet + " RETURN position"

	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func DeletePosition(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	positionId := params["positionId"]
	query := fmt.Sprintf("MATCH (position:Position) WHERE ID(position)=%v DETACH DELETE position", positionId)
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

func CreatePosition(w http.ResponseWriter, r *http.Request) {
	var position models.Position
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&position); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	query := fmt.Sprintf("CREATE (position:Position {name: \"%v\"})", position.Name)

	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
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

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	params := mux.Vars(r)
	projectId := params["projectId"]
	query := fmt.Sprintf("MATCH (project:Project) WHERE ID(project)=%v DETACH DELETE project", projectId)
	_, err := session.Run(query, nil)
	if err != nil {
		fmt.Errorf("not found: %v", projectId)
		fmt.Errorf("message: %v", err)
		respondWithError(w, 400, err.Error())
	} else {
		respondWithJSON(w, 200, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectId := params["projectId"]
	queryMatch := fmt.Sprintf("MATCH (project:Project) WHERE ID(project)=%v", projectId)
	querySet := " SET"

	var project models.Project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if project.Name != "" {
		querySet += fmt.Sprintf(" project.name = \"%v\"", project.Name)
	}

	query := queryMatch + querySet + " RETURN project"

	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	query := fmt.Sprintf("CREATE (project:Project {name: \"%v\"})", project.Name)

	_, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func CreateWorksAs(w http.ResponseWriter, r *http.Request) {
	var worksAs models.WorksAs
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&worksAs); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	query := fmt.Sprintf("MATCH  (employee:Employee),  (position:Position) WHERE ID(employee) = %v AND ID(position) = %v AND NOT (employee)-[:WorksAs]->(position) CREATE (employee)-[r:WorksAs]->(position) RETURN type(r)", worksAs.EmployeeId, worksAs.PositionId)
	result, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	record, err := result.Single()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	relType, _ := record.Get("type(r)")
	if err != nil && relType == "WorksAs" {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
}

func CreateWorksIn(w http.ResponseWriter, r *http.Request) {
	var worksIn models.WorksIn
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&worksIn); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	session := (*DbDriver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	query := fmt.Sprintf("MATCH  (employee:Employee),  (project:Project) WHERE ID(employee) = %v AND ID(project) = %v AND NOT (employee)-[:WorksIn]->(project) CREATE (employee)-[r:WorksIn]->(project) RETURN type(r)", worksIn.EmployeeId, worksIn.ProjectId)
	result, err := session.Run(query, nil)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	record, err := result.Single()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	relType, _ := record.Get("type(r)")
	if err != nil && relType == "WorksIn" {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		respondWithJSON(w, http.StatusCreated, models.ModificationStatus{Status: "ok", Error: ""})
	}
}
