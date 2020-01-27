package multilock

import (
	"sync"
)

type Var struct {
	length       int
	mutexes      []sync.Mutex
	distribution func(s string, length int) int
	global       sync.Mutex
}

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

func (m *Var) Get(i interface{}) *sync.Mutex {
	return m.GetID(addr(i))
}

func (m *Var) GetID(id string) *sync.Mutex {
	m.global.Lock()
	defer m.global.Unlock()

	index := m.distribution(id, m.length)
	return &m.mutexes[index]
}

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
