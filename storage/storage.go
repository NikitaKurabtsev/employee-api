package storage

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary string `json:"salary"`
}
