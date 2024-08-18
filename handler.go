package main

import (
	"log/slog"
	"net/http"

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
