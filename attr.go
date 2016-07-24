package attr

import "sync"

// Entity is object
type Entity struct {
	mu   sync.Mutex
	data map[interface{}]interface{}
}

// NewEntity create entity
func NewEntity() *Entity {
	return &Entity{
		data: map[interface{}]interface{}{},
	}
}

func (e *Entity) get(key interface{}) (interface{}, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()

	d, ok := e.data[key]
	if !ok {
		return nil, false
	}

	return d, true
}

func (e *Entity) set(key interface{}, value interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

}

// Exists check the key existence
func (e *Entity) Exists(key interface{}) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	_, ok := e.data[key]
	return ok
}
