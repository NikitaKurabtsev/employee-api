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

type ErrorResponse struct {
	Message string `json:"message"`
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
		h.logger.Error("Failed to bind JSON", "json", err)
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
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
		h.logger.Error("Failed to convert id param to int",
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusNotFound, employee)
		return
	}

	c.JSON(http.StatusOK, employee)
}
