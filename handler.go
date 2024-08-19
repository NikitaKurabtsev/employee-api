package main

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	List() []Employee
	Update(id int, e *Employee) error
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
	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		RespondWithError(c, h.logger, http.StatusBadRequest, "CreateEmployee: Failed to bind JSON", err)
		return
	}

	if err := validateEmployee(employee); err != nil {
		RespondWithError(c, h.logger, http.StatusBadRequest, "CreateEmployee: invalid employee data", err)
		return
	}

	h.storage.Insert(&employee)
	h.logger.Info("Employee created:", "name", employee.Name, "id", employee.Id)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": employee.Id,
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RespondWithError(c, h.logger, http.StatusBadRequest, "GetEmployee: invalid ID", err)
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		RespondWithError(c, h.logger, http.StatusNotFound, "GetEmployee: employee not found", err)
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		RespondWithError(c, h.logger, http.StatusBadRequest, "UpdateEmployee: invalid ID", err)
		return
	}

	var employee Employee

	if err := c.BindJSON(&employee); err != nil {
		RespondWithError(c, h.logger, http.StatusBadRequest, "UpdateEmployee: failed to parse JSON", err)
		return
	}

	if err := validateEmployee(employee); err != nil {
		RespondWithError(c, h.logger, http.StatusBadRequest, "UpdateEmployee: invalid employee data", err)
		return
	}

	if err := h.storage.Update(id, &employee); err != nil {
		RespondWithError(c, h.logger, http.StatusBadRequest, "UpdateEmployee: failed to update employee", err)
		return
	}

	c.JSON(http.StatusOK, employee)
}
