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
}

func (s *MapMemoryStorage) Get(id int) (Employee, error) {
	e, ok := s.data[id]
	if !ok {
		return Employee{}, errors.New("employee with such id doesn't exists")
	}
	return e, nil
}
