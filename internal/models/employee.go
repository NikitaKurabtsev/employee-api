package models

type Employee struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Sex         string  `json:"sex"`
	Age         int     `json:"age"`
	Salary      float32 `json:"salary"`
	PhoneNumber string  `json:"phone_number"`
}
