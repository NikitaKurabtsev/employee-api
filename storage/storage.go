package storage

import (
	"errors"
	"sync"
)

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary string `json:"salary"`
}

type MapMemoryStorage struct {
	counter int
	data    map[int]Employee
	mu      sync.Mutex
}

func NewMapMemoryStorage() *MapMemoryStorage {
	return &MapMemoryStorage{
		counter: 1,
		data:    make(map[int]Employee),
	}
}

func (s *MapMemoryStorage) Insert(e *Employee) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.Id = s.counter
	s.data[e.Id] = *e
	s.counter++
}

func (s *MapMemoryStorage) Get(id int) (Employee, error) {
	if emp, ok := s.data[id]; ok {
		return emp, nil
	}
	return Employee{}, errors.New("employee with such id doesn't exists")
}

func (s *MapMemoryStorage) List() []Employee {
	var employeeList []Employee
	for _, e := range s.data {
		employeeList = append(employeeList, e)
	}
	return employeeList
}

func (s *MapMemoryStorage) Update(id int, e *Employee) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; ok {
		s.data[id] = *e
		return *e, nil
	}
	return *e, errors.New("failed to update")
}

func (s *MapMemoryStorage) Delete(id int) error {
	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return nil
	}
	return errors.New("employee with such id doesn't exists")
}
