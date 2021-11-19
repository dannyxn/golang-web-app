package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {

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

func createEmployee(w http.ResponseWriter, r *http.Request) {
}

func readEmployee(w http.ResponseWriter, r *http.Request) {
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
