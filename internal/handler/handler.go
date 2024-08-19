package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/NikitaKurabtsev/employee-api.git/internal/errors"
	"github.com/NikitaKurabtsev/employee-api.git/internal/models"
	"github.com/NikitaKurabtsev/employee-api.git/internal/validation"
	"github.com/gin-gonic/gin"
)

const (
	createMethodName = "CreateEmployee"
	updateMethodName = "UpdateEmployee"
	getAllMethodName = "GetAllEmployees"
	getMethodName    = "GetEmployee"
	delereMethodName = "DeleteEmployee"
)

type Repository interface {
	Insert(e *models.Employee)
	Get(id int) (models.Employee, error)
	List() []models.Employee
	Update(id int, e *models.Employee) error
	Delete(id int) error
}

type Handler struct {
	repository Repository
	logger     *slog.Logger
}

// NewHandler returns pointer to the Handler
// and implements Dependency Injection pattern
func NewHandler(repository Repository, logger *slog.Logger) *Handler {
	return &Handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee models.Employee
	var errMessage string

	if err := c.BindJSON(&employee); err != nil {
		errMessage = errors.ErrMessage(createMethodName, errors.ErrInvalidJSON)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	if err := validation.ValidateEmployee(employee); err != nil {
		errMessage = errors.ErrMessage(createMethodName, errors.ErrEmployeeValidation)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	h.repository.Insert(&employee)
	h.logger.Info("employee created", "name", employee.Name, "id", employee.Id)

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created",
		"id":      employee.Id,
	})
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	allEmployees := h.repository.List()
	count := len(allEmployees)

	c.JSON(http.StatusOK, gin.H{
		"employees": allEmployees,
		"count":     count,
	})
}

func (h *Handler) GetEmployee(c *gin.Context) {
	var errMessage string

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMessage = errors.ErrMessage(getMethodName, errors.ErrEmployeeInvalidId)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	employee, err := h.repository.Get(id)
	if err != nil {
		errMessage = errors.ErrMessage(getMethodName, errors.ErrEmployeeNotFound)
		errors.RespondWithError(c, h.logger, http.StatusNotFound, errMessage, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	var errMessage string

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMessage = errors.ErrMessage(updateMethodName, errors.ErrEmployeeInvalidId)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		errMessage = errors.ErrMessage(updateMethodName, errors.ErrInvalidJSON)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	if err := validation.ValidateEmployee(employee); err != nil {
		errMessage = errors.ErrMessage(updateMethodName, errors.ErrEmployeeValidation)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	if err := h.repository.Update(id, &employee); err != nil {
		errMessage = errors.ErrMessage(updateMethodName, errors.ErrEmployeeUpdate)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "updated successfully",
		"employee": employee,
	})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	var errMessage string

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMessage = errors.ErrMessage(delereMethodName, errors.ErrEmployeeInvalidId)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	err = h.repository.Delete(id)
	if err != nil {
		errors.RespondWithError(c, h.logger, http.StatusNotFound, "DeleteEmployee: employee not found", err)
		return
	}

	h.logger.Info("DeleteEmployee: employee deleted", "id", id)

	c.Status(http.StatusNoContent)
}
