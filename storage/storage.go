package storage

import "sync"

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

func (m *MapMemoryStorage) Insert(e *Employee) {
	m.mu.Lock()
	defer m.mu.Unlock()

	e.Id = m.counter
	m.data[e.Id] = *e
}
