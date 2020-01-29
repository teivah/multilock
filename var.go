package multilock

import (
	"sync"
)

// Var is a variable length structure of sync.Mutex
type Var struct {
	length       int
	mutexes      []sync.Mutex
	distribution func(i interface{}, length int) int
	global       sync.Mutex
}

// NewVar creates a variable length structure of sync.Mutex
func NewVar(length int, opts ...Option) *Var {
	var options options
	for _, opt := range opts {
		opt.apply(&options)
	}

	if options.distribution == nil {
		options.distribution = distribution
	}

	return &Var{
		length:       length,
		distribution: options.distribution,
		mutexes:      make([]sync.Mutex, length),
	}
}

// Get retrieves a sync.Mutex from an interface
func (m *Var) Get(i interface{}) *sync.Mutex {
	return m.GetID(addr(i))
}

// GetID retrieves a sync.Mutex from an identifier
func (m *Var) GetID(id string) *sync.Mutex {
	m.global.Lock()
	defer m.global.Unlock()

	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}

// Resize the internal multilock structure
func (m *Var) Resize(length int, opts ...Option) {
	m.global.Lock()
	defer m.global.Unlock()

	var options options
	for _, opt := range opts {
		opt.apply(&options)
	}

	if options.distribution != nil {
		m.distribution = options.distribution
	}

	m.mutexes = make([]sync.Mutex, length)
}
