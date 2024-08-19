package repository

import (
	"fmt"
	"sync"

	"github.com/NikitaKurabtsev/employee-api.git/internal/models"
)

const (
	errNotFound       = "employee doesn't exists"
	errFailedToUpdate = "failed to update employee"
)

type EmployeeRepository struct {
	counter int
	data    map[int]models.Employee
	mu      sync.Mutex
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		counter: 1,
		data:    make(map[int]models.Employee),
	}
}

func (r *EmployeeRepository) Insert(e *models.Employee) {
	r.mu.Lock()
	defer r.mu.Unlock()

	e.Id = r.counter
	r.data[e.Id] = *e
	r.counter++
}

func (r *EmployeeRepository) Get(id int) (models.Employee, error) {
	if emp, ok := r.data[id]; ok {
		return emp, nil
	}
	return models.Employee{}, fmt.Errorf(errNotFound)
}

func (r *EmployeeRepository) List() []models.Employee {
	var employeeList []models.Employee

	for _, e := range r.data {
		employeeList = append(employeeList, e)
	}
	return employeeList
}

func (r *EmployeeRepository) Update(id int, e *models.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; ok {
		r.data[id] = *e
		return nil
	}
	return fmt.Errorf(errFailedToUpdate)
}

func (r *EmployeeRepository) Delete(id int) error {
	if _, ok := r.data[id]; ok {
		delete(r.data, id)
		return nil
	}
	return fmt.Errorf(errNotFound)
}
