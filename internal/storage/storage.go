package storage

import (
	"fmt"
	"sync"

	"github.com/NikitaKurabtsev/employee-api.git/internal/models"
)

var (
	errEmployeeNotFound       = fmt.Errorf("employee doesn't exists")
	errFailedToUpdateEmployee = fmt.Errorf("failed to update employee")
)

type MapMemoryStorage struct {
	counter int
	data    map[int]models.Employee
	mu      sync.Mutex
}

func NewMapMemoryStorage() *MapMemoryStorage {
	return &MapMemoryStorage{
		counter: 1,
		data:    make(map[int]models.Employee),
	}
}

func (s *MapMemoryStorage) Insert(e *models.Employee) {
	s.mu.Lock()
	defer s.mu.Unlock()

	e.Id = s.counter
	s.data[e.Id] = *e
	s.counter++
}

func (s *MapMemoryStorage) Get(id int) (models.Employee, error) {
	if emp, ok := s.data[id]; ok {
		return emp, nil
	}
	return models.Employee{}, errEmployeeNotFound
}

func (s *MapMemoryStorage) List() []models.Employee {
	var employeeList []models.Employee

	for _, e := range s.data {
		employeeList = append(employeeList, e)
	}
	return employeeList
}

func (s *MapMemoryStorage) Update(id int, e *models.Employee) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[id]; ok {
		s.data[id] = *e
		return nil
	}
	return errFailedToUpdateEmployee
}

func (s *MapMemoryStorage) Delete(id int) error {
	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return nil
	}
	return errEmployeeNotFound
}
