package db

import (
	"context"
	"sync"

	"github.com/110y/tododemo/internal/todo"
)

var _ DB = (*memoryDB)(nil)

type memoryDB struct {
	db   map[string]*todo.TODO
	lock sync.RWMutex
}

func NewMemoryDB() DB {
	return &memoryDB{db: map[string]*todo.TODO{}}
}

func (m *memoryDB) GetAllTODOs(ctx context.Context) ([]*todo.TODO, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make([]*todo.TODO, len(m.db))
	i := 0
	for _, t := range m.db {
		result[i] = t
		i++
	}

	return result, nil
}

func (m *memoryDB) PutTODO(ctx context.Context, t *todo.TODO) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()

	return nil
}
