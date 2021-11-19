package main

type Employee struct {
	Name        string `json:"Name"`
	Surname     string `json:"Surname"`
	PositionID  string `json:"PositionID"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Position struct {
	Name string `json:"Name"`
}
