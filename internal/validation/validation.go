package validation

import (
	"fmt"

	"github.com/NikitaKurabtsev/employee-api/internal/models"
)

type EmployeeValidator struct{}

func NewEmployeeValidator() *EmployeeValidator {
	return &EmployeeValidator{}
}

func (v EmployeeValidator) Validate(e models.Employee) error {
	if e.Name == "" || e.Age < 0 || e.Salary < 0 {
		return fmt.Errorf("invalid employee data")
	}
	return nil
}
