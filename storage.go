package main

import (
	"fmt"
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
	return Employee{}, fmt.Errorf("employee with id: %d doesn't exists", id)
}

func (s *MapMemoryStorage) List() []Employee {
	var employeeList []Employee
	for _, e := range s.data {
		employeeList = append(employeeList, e)
	}
	return employeeList
}

func (s *MapMemoryStorage) Update(id int, e *Employee) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; ok {
		s.data[id] = *e
		return nil
	}
	return fmt.Errorf("failed to update the employee with the id: %d", id)
}

func (s *MapMemoryStorage) Delete(id int) error {
	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return nil
	}
	return fmt.Errorf("employee with id: %d doesn't exists", id)
}
