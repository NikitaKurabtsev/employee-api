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

type Storage interface {
	Insert(e *models.Employee)
	Get(id int) (models.Employee, error)
	List() []models.Employee
	Update(id int, e *models.Employee) error
	Delete(id int) error
}

type Handler struct {
	storage Storage
	logger  *slog.Logger
}

// NewHandler returns pointer to the Handler
// and implements Dependency Injection pattern
func NewHandler(storage Storage, logger *slog.Logger) *Handler {
	return &Handler{
		storage: storage,
		logger:  logger,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	funcName := "CreateEmployee"
	var employee models.Employee
	var errMessage string

	if err := c.BindJSON(&employee); err != nil {
		errMessage = errors.ErrMessage(funcName, errors.ErrInvalidJSON)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	if err := validation.ValidateEmployee(employee); err != nil {
		errMessage = errors.ErrMessage(funcName, errors.ErrEmployeeValidation)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	h.storage.Insert(&employee)
	h.logger.Info("employee created", "name", employee.Name, "id", employee.Id)

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created",
		"id":      employee.Id,
	})
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	allEmployees := h.storage.List()
	count := len(allEmployees)

	c.JSON(http.StatusOK, gin.H{
		"employees": allEmployees,
		"count":     count,
	})
}

func (h *Handler) GetEmployee(c *gin.Context) {
	funcName := "GetEmployee"
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMessage := errors.ErrMessage(funcName, errors.ErrInvalidId)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		errMessage := errors.ErrMessage(funcName, errors.ErrEmployeeNotFound)
		errors.RespondWithError(c, h.logger, http.StatusNotFound, errMessage, err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMessage := errors.ErrMessage("UpdateEmployee:", errors.ErrInvalidId)
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, errMessage, err)
		return
	}

	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, "UpdateEmployee: failed to parse JSON", err)
		return
	}

	if err := validation.ValidateEmployee(employee); err != nil {
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, "UpdateEmployee: invalid employee data", err)
		return
	}

	if err := h.storage.Update(id, &employee); err != nil {
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, "UpdateEmployee: failed to update employee", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "updated successfully",
		"employee": employee,
	})
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.RespondWithError(c, h.logger, http.StatusBadRequest, "DeleteEmployee: invalid ID", err)
		return
	}

	err = h.storage.Delete(id)
	if err != nil {
		errors.RespondWithError(c, h.logger, http.StatusNotFound, "DeleteEmployee: employee not found", err)
		return
	}

	h.logger.Info("DeleteEmployee: employee deleted", "id", id)

	c.Status(http.StatusNoContent)
}
