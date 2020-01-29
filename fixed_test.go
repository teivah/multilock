package multilock

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixed_Get(t *testing.T) {
	s := []int{1}
	fixed := NewFixed(10)
	m := fixed.Get(s)
	m.Lock()
	defer m.Unlock()
	assert.Equal(t, addr(fixed.Get(s)), addr(m))
}

func TestFixed_Get_One(t *testing.T) {
	s := []int{1}
	s2 := []int{1, 2}
	fixed := NewFixed(1)
	m := fixed.Get(s)
	m.Lock()
	defer m.Unlock()
	assert.Equal(t, addr(fixed.Get(s2)), addr(m))
}

func TestFixed_Get_CustomDistribution(t *testing.T) {
	s := []int{1}
	s2 := []int{1, 2}
	fixed := NewFixed(100, WithCustomDistribution(func(_ interface{}, _ int) int {
		return 0
	}))
	m := fixed.Get(s)
	m.Lock()
	defer m.Unlock()
	assert.Equal(t, addr(fixed.Get(s2)), addr(m))
}

func TestFixed_Get_Race(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	fixed := NewFixed(10)

	go func() {
		m := fixed.Get("")
		m.Lock()
		defer m.Unlock()
		wg.Done()
	}()

	go func() {
		m := fixed.Get("")
		m.Lock()
		defer m.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
