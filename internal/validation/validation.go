package validation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/NikitaKurabtsev/employee-api/internal/models"
)

type EmployeeValidator struct{}

func NewEmployeeValidator() *EmployeeValidator {
	return &EmployeeValidator{}
}

func (v EmployeeValidator) ValidateFields(e models.Employee) error {
	if e.Name == "" || e.Age < 0 || e.Salary < 0 {
		return fmt.Errorf("invalid employee data")
	}
	return nil
}

func (v EmployeeValidator) ValidateId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}
