package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/NikitaKurabtsev/employee-api/internal/errors"
	"github.com/NikitaKurabtsev/employee-api/internal/httputils"
	"github.com/NikitaKurabtsev/employee-api/internal/models"
	"github.com/NikitaKurabtsev/employee-api/internal/validation"
	"github.com/gin-gonic/gin"
)

const (
	// method names
	createMethodName = "CreateEmployee"
	updateMethodName = "UpdateEmployee"
	getAllMethodName = "GetAllEmployees"
	getMethodName    = "GetEmployee"
	deleteMethodName = "DeleteEmployee"

	// success messages
	okCreated = "employee created"
	okUpdated = "employee updated"
	okDeleted = "employee deleted"
	okReaded  = "employee readed"
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

	if err := c.BindJSON(&employee); err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(createMethodName, errors.ErrInvalidJSON),
			err,
		)
		return
	}

	if err := validation.ValidateEmployee(employee); err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(createMethodName, errors.ErrEmployeeValidation),
			err,
		)
		return
	}

	h.repository.Insert(&employee)

	httputils.RespondWithStatus(
		c,
		h.logger,
		http.StatusCreated,
		httputils.OkMessage(createMethodName, okCreated),
		employee,
	)
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	allEmployees := h.repository.List()

	httputils.RespondWithStatus(
		c,
		h.logger,
		http.StatusOK,
		httputils.OkMessage(getAllMethodName, okReaded),
		allEmployees,
	)
}

func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(getMethodName, errors.ErrEmployeeInvalidId),
			err,
		)
		return
	}

	employee, err := h.repository.Get(id)
	if err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusNotFound,
			errors.ErrMessage(getMethodName, errors.ErrEmployeeNotFound),
			err,
		)
		return
	}

	httputils.RespondWithStatus(
		c,
		h.logger,
		http.StatusOK,
		httputils.OkMessage(getMethodName, okReaded),
		employee,
	)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(updateMethodName, errors.ErrEmployeeInvalidId),
			err,
		)
		return
	}

	var employee models.Employee
	if err := c.BindJSON(&employee); err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(updateMethodName, errors.ErrInvalidJSON),
			err,
		)
		return
	}
	employee.Id = id

	if err := validation.ValidateEmployee(employee); err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(updateMethodName, errors.ErrEmployeeValidation),
			err,
		)
		return
	}

	if err := h.repository.Update(id, &employee); err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(updateMethodName, errors.ErrEmployeeUpdate),
			err,
		)
		return
	}

	httputils.RespondWithStatus(
		c,
		h.logger,
		http.StatusOK,
		httputils.OkMessage(updateMethodName, okUpdated),
		employee,
	)
}

func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusBadRequest,
			errors.ErrMessage(deleteMethodName, errors.ErrEmployeeInvalidId),
			err,
		)
		return
	}

	err = h.repository.Delete(id)
	if err != nil {
		errors.RespondWithError(
			c,
			h.logger,
			http.StatusNotFound,
			errors.ErrMessage(deleteMethodName, errors.ErrEmployeeNotFound),
			err,
		)
		return
	}

	httputils.RespondWithStatus(
		c,
		h.logger,
		http.StatusNoContent,
		httputils.OkMessage(deleteMethodName, okDeleted),
		0,
	)
}
