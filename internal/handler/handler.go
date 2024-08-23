package handler

import (
	"github.com/NikitaKurabtsev/employee-api/internal/interfaces"
	"github.com/NikitaKurabtsev/employee-api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrEmployeeInvalidId  = "ID must be a number"
	ErrInvalidJSON        = "failed to parse JSON"
	ErrEmployeeUpdate     = "failed to update employee"
	ErrEmployeeNotFound   = "employee not found"
	ErrEmployeeValidation = "invalid employee data"
)

type (
	Repository interface {
		Insert(e *models.Employee)
		Get(id int) (models.Employee, error)
		List() []models.Employee
		Update(id int, e *models.Employee) error
		Delete(id int) error
	}
	Responder interface {
		RespondWithStatus(c *gin.Context, statusCode int, data any)
		RespondWithError(c *gin.Context, logger interfaces.Logger, statusCode int, customError string, originalErr error)
	}
	Validator interface {
		ValidateFields(e models.Employee) error
		ValidateId(c *gin.Context) (int, error)
	}
)

type Handler struct {
	repository Repository
	respond    Responder
	logger     interfaces.Logger
	validation Validator
}

// NewHandler returns pointer to the Handler
// and implements Dependency Injection pattern
func NewHandler(
	repository Repository,
	respond Responder,
	logger interfaces.Logger,
	validation Validator,
) *Handler {
	return &Handler{
		repository: repository,
		respond:    respond,
		logger:     logger,
		validation: validation,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrInvalidJSON, err)
		return
	}

	if err := h.validation.ValidateFields(employee); err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrEmployeeValidation, err)
		return
	}

	h.repository.Insert(&employee)

	h.respond.RespondWithStatus(c, http.StatusCreated, employee)
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	allEmployees := h.repository.List()

	h.respond.RespondWithStatus(c, http.StatusOK, allEmployees)
}

func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := h.validation.ValidateId(c)
	if err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrEmployeeInvalidId, err)
		return
	}

	employee, err := h.repository.Get(id)
	if err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusNotFound, ErrEmployeeNotFound, err)
		return
	}

	h.respond.RespondWithStatus(c, http.StatusOK, employee)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := h.validation.ValidateId(c)
	if err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrEmployeeInvalidId, err)
		return
	}

	var employee models.Employee
	if err = c.BindJSON(&employee); err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrInvalidJSON, err)
		return
	}
	employee.Id = id

	if err = h.validation.ValidateFields(employee); err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrEmployeeValidation, err)
		return
	}

	if err = h.repository.Update(id, &employee); err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrEmployeeUpdate, err)
		return
	}

	h.respond.RespondWithStatus(c, http.StatusOK, employee)
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := h.validation.ValidateId(c)
	if err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusBadRequest, ErrEmployeeInvalidId, err)
		return
	}

	err = h.repository.Delete(id)
	if err != nil {
		h.respond.RespondWithError(c, h.logger, http.StatusNotFound, ErrEmployeeNotFound, err)
		return
	}

	h.respond.RespondWithStatus(c, http.StatusNoContent, struct{}{})
}
