package model

type Employee struct {
	Id     string  `json:"id,omitempty" bson:"_id"`
	Name   string  `json:"name,omitempty"`
	Email  string  `json:"email,omitempty"`
	Salary float64 `json:"salary,omitempty"`
}
