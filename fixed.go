package multilock

import (
	"sync"
)

// Fixed is a fixed length structure of sync.Mutex
type Fixed struct {
	length       int
	mutexes      []sync.Mutex
	distribution func(i interface{}, length int) int
}

// NewFixed creates a fixed length structure of sync.Mutex
func NewFixed(length int, opts ...Option) *Fixed {
	var options options
	for _, opt := range opts {
		opt.apply(&options)
	}

	if options.distribution == nil {
		options.distribution = distribution
	}

	return &Fixed{
		length:       length,
		distribution: options.distribution,
		mutexes:      make([]sync.Mutex, length),
	}
}

// Get retrieves a sync.Mutex from an interface
func (m *Fixed) Get(i interface{}) *sync.Mutex {
	return m.GetID(addr(i))
}

// GetID retrieves a sync.Mutex from an identifier
func (m *Fixed) GetID(id string) *sync.Mutex {
	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}
