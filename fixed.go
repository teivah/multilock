package multilock

import (
	"sync"
)

type Fixed struct {
	length       int
	mutexes      []sync.Mutex
	distribution func(s string, length int) int
}

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

func (m *Fixed) Get(i interface{}) *sync.Mutex {
	return m.GetID(addr(i))
}

func (m *Fixed) GetID(id string) *sync.Mutex {
	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}
