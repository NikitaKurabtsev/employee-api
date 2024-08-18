package main

import (
	"log/slog"
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
