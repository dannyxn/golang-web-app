package models

type Employee struct {
	Id          int64  `json:"Id"`
	Name        string `json:"Name"`
	Surname     string `json:"Surname"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Position struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

type Project struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

type ModificationStatus struct {
	Status string `json:"Status"`
	Error  string `json:"Error"`
}
