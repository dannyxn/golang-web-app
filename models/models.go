package models

type Employee struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Surname     string `json:"Surname"`
	PositionID  string `json:"PositionID"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Position struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type Project struct {
	ID           string    `json:"ID"`
	Name         string    `json:"Name"`
	ProjectOwner *Employee `json:"ProjectOwner"`
}
