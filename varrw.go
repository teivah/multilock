package multilock

import (
	"sync"
)

type VarRW struct {
	length       int
	mutexes      []sync.RWMutex
	distribution func(s string, length int) int
	global       sync.Mutex
}

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

func (m *VarRW) Get(i interface{}) *sync.RWMutex {
	return m.GetID(addr(i))
}

func (m *VarRW) GetID(id string) *sync.RWMutex {
	m.global.Lock()
	defer m.global.Unlock()

	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}

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
