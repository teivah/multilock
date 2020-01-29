package multilock

import (
	"sync"
)

// FixedRW is a fixed length structure of sync.RWMutex
type FixedRW struct {
	length       int
	mutexes      []sync.RWMutex
	distribution func(i interface{}, length int) int
}

// NewFixedRW creates a fixed length structure of sync.RWMutex
func NewFixedRW(length int, opts ...Option) *FixedRW {
	var options options
	for _, opt := range opts {
		opt.apply(&options)
	}

	if options.distribution == nil {
		options.distribution = distribution
	}

	return &FixedRW{
		length:       length,
		distribution: options.distribution,
		mutexes:      make([]sync.RWMutex, length),
	}
}

// Get retrieves a sync.RXMutex from an interface
func (m *FixedRW) Get(i interface{}) *sync.RWMutex {
	return m.GetID(addr(i))
}

// GetID retrieves a sync.RWMutex from an identifier
func (m *FixedRW) GetID(id string) *sync.RWMutex {
	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}
