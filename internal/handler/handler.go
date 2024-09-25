package handler

import (
	"github.com/NikitaKurabtsev/employee-api/internal/models"
	"github.com/NikitaKurabtsev/employee-api/internal/validation"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const (
	ErrEmployeeInvalidID  = "ID must be a number"
	ErrInvalidJSON        = "failed to parse JSON"
	ErrEmployeeUpdate     = "failed to update employee"
	ErrEmployeeNotFound   = "employee not found"
	ErrEmployeeValidation = "invalid employee data"
)

const (
	createHandler = "CreateEmployee"
	getHandler    = "GetEmployee"
	updateHandler = "UpdateEmployee"
	deleteHandler = "DeleteEmployee"
)

type Repository interface {
	Insert(e models.Employee)
	Get(id int) (models.Employee, error)
	List() []models.Employee
	Update(id int, e models.Employee) error
	Delete(id int) error
}

type Handler struct {
	repository Repository
	logger     *slog.Logger
}

// NewHandler returns pointer to the Handler
// and implements Dependency Injection pattern
func NewHandler(
	repository Repository,
	logger *slog.Logger,
) *Handler {
	return &Handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		h.logger.Error("%s: %s", createHandler, ErrInvalidJSON)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validation.ValidateFields(employee); err != nil {
		h.logger.Error("%s: %s", createHandler, ErrEmployeeValidation)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.repository.Insert(employee)

	c.JSON(http.StatusCreated, gin.H{"created": employee})
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	allEmployees := h.repository.List()

	c.JSON(http.StatusOK, allEmployees)
}

func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := validation.ValidateId(c.Param("id"))
	if err != nil {
		h.logger.Error("%s: %s", getHandler, ErrEmployeeInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := h.repository.Get(id)
	if err != nil {
		h.logger.Error("%s: %s", getHandler, ErrEmployeeNotFound)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := validation.ValidateId(c.Param("id"))
	if err != nil {
		h.logger.Error("%s: %s", updateHandler, ErrEmployeeInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var employee models.Employee
	if err = c.BindJSON(&employee); err != nil {
		h.logger.Error("%s: %s", updateHandler, ErrInvalidJSON)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.Id = id

	if err = validation.ValidateFields(employee); err != nil {
		h.logger.Error("%s: %s", updateHandler, ErrEmployeeValidation)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.repository.Update(id, employee); err != nil {
		h.logger.Error("%s: %s", updateHandler, ErrEmployeeUpdate)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": employee})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := validation.ValidateId(c.Param("id"))
	if err != nil {
		h.logger.Error("%s: %s", deleteHandler, ErrEmployeeInvalidID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.repository.Delete(id)
	if err != nil {
		h.logger.Error("%s: %s", deleteHandler, ErrEmployeeNotFound)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
