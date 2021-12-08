package models

// nodes

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

// relationships

type WorksAs struct {
	EmployeeId int64 `json:"EmployeeId"`
	PositionId int64 `json:"PositionId"`
}

type WorksIn struct {
	EmployeeId int64 `json:"EmployeeId"`
	ProjectId  int64 `json:"ProjectId"`
}

type ModificationStatus struct {
	Status string `json:"Status"`
	Error  string `json:"Error"`
}
