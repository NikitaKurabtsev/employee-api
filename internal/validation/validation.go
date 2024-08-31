package validation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"

	"github.com/NikitaKurabtsev/employee-api/internal/models"
)

type EmployeeValidator struct{}

func NewEmployeeValidator() *EmployeeValidator {
	return &EmployeeValidator{}
}

func (v EmployeeValidator) ValidateFields(e models.Employee) error {
	pattern := `(\+7|8)?\d{10}$`
	re := regexp.MustCompile(pattern)
	isPhoneValid := re.MatchString(e.PhoneNumber)

	if !isPhoneValid {
		return fmt.Errorf("invalid phone number")
	}
	if e.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if e.Age < 0 {
		return fmt.Errorf("age cannot be negative")
	}
	if e.Salary < 0 {
		return fmt.Errorf("salary cannot be negative")
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
