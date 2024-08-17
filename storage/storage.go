package storage

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary string `json:"salary"`
}

func NewEmployee(name, sex, salary string, age int) *Employee {
	return &Employee{
		Name:   name,
		Sex:    sex,
		Age:    age,
		Salary: salary,
	}
}
