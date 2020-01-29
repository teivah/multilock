package multilock

import (
	"sync"
)

// VarRW is a variable length structure of sync.RWMutex
type VarRW struct {
	length       int
	mutexes      []sync.RWMutex
	distribution func(i interface{}, length int) int
	global       sync.Mutex
}

// NewVarRW creates a variable length structure of sync.RWMutex
func NewVarRW(length int, opts ...Option) *VarRW {
	var options options
	for _, opt := range opts {
		opt.apply(&options)
	}

	if options.distribution == nil {
		options.distribution = distribution
	}

	return &VarRW{
		length:       length,
		distribution: options.distribution,
		mutexes:      make([]sync.RWMutex, length),
	}
}

// Get retrieves a sync.RWMutex from an interface
func (m *VarRW) Get(i interface{}) *sync.RWMutex {
	return m.GetID(addr(i))
}

// GetID retrieves a sync.RWMutex from an identifier
func (m *VarRW) GetID(id string) *sync.RWMutex {
	m.global.Lock()
	defer m.global.Unlock()

	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}

// Resize the internal multilock structure
func (m *VarRW) Resize(length int, opts ...Option) {
	m.global.Lock()
	defer m.global.Unlock()

	var options options
	for _, opt := range opts {
		opt.apply(&options)
	}

	if options.distribution != nil {
		m.distribution = options.distribution
	}

	m.mutexes = make([]sync.RWMutex, length)
}
