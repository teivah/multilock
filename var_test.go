package multilock

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVar_Get(t *testing.T) {
	s := []int{1}
	v := NewVar(10)
	m := v.Get(s)
	m.Lock()
	defer m.Unlock()
	assert.Equal(t, addr(v.Get(s)), addr(m))
}

func TestVar_Get_One(t *testing.T) {
	s := []int{1}
	s2 := []int{1, 2}
	v := NewVar(1)
	m := v.Get(s)
	m.Lock()
	defer m.Unlock()
	assert.Equal(t, addr(v.Get(s2)), addr(m))
}

func TestVar_Get_CustomDistribution(t *testing.T) {
	s := []int{1}
	s2 := []int{1, 2}
	v := NewVar(100, WithCustomDistribution(func(_ interface{}, _ int) int {
		return 0
	}))
	m := v.Get(s)
	m.Lock()
	defer m.Unlock()
	assert.Equal(t, addr(v.Get(s2)), addr(m))
}

func TestVar_Get_Race(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	v := NewVar(10)

	go func() {
		m := v.Get("")
		m.Lock()
		defer m.Unlock()
		wg.Done()
		v.Resize(20)
	}()

	go func() {
		m := v.Get("")
		m.Lock()
		defer m.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
