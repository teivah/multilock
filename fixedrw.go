package multilock

import (
	"sync"
)

type FixedRW struct {
	length       int
	mutexes      []sync.RWMutex
	distribution func(s string, length int) int
}

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

func (m *FixedRW) Get(i interface{}) *sync.RWMutex {
	return m.GetID(addr(i))
}

func (m *FixedRW) GetID(id string) *sync.RWMutex {
	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}
