package validation

import (
	"fmt"

	"github.com/NikitaKurabtsev/employee-api/internal/models"
)

func ValidateEmployee(e models.Employee) error {
	if e.Name == "" || e.Age < 0 || e.Salary < 0 {
		return fmt.Errorf("name and age must be valid")
	}
	return nil
}
