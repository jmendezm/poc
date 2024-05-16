package memory_db

import (
	"sync"
	"technical_test/domain"
)

var connections_memory_db = &memDB{
	mu:          sync.RWMutex{},
	connections: make(map[string]*domain.Connection),
}

type memDB struct {
	mu          sync.RWMutex
	connections map[string]*domain.Connection
}

func GetConnectionByID(id string) *domain.Connection {
	if connections_memory_db == nil {
		initMemDB()
	}
	connections_memory_db.mu.RLock()
	defer connections_memory_db.mu.RUnlock()
	if c, ok := connections_memory_db.connections[id]; ok {
		return c
	}
	return nil
}

func SetConnection(id string, conn *domain.Connection) {
	if connections_memory_db == nil {
		initMemDB()
	}
	connections_memory_db.mu.Lock()
	defer connections_memory_db.mu.Unlock()
	connections_memory_db.connections[id] = conn
}

func initMemDB() {
	connections_memory_db = &memDB{
		mu:          sync.RWMutex{},
		connections: make(map[string]*domain.Connection),
	}
}
