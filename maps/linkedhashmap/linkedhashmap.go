package linkedhashmap

import (
	"errors"

	"github.com/prprprus/ds/list/doublelinkedlist"
)

var (
	// ErrNotFound is returned when the value can not found
	ErrNotFound = errors.New("can not found value")
	// ErrEmpty is returned when the map is empty
	ErrEmpty = errors.New("the map is empty")
)

type Map struct {
	m        map[interface{}]interface{}
	ordering *doublelinkedlist.List
}

func New() *Map {
	return &Map{
		m:        make(map[interface{}]interface{}),
		ordering: doublelinkedlist.New(),
	}
}

// Map Interface

func (m *Map) Put(key, value interface{}) {
	if _, ok := m.m[key]; !ok {
		m.ordering.Append(value)
	}
	m.m[key] = value
}

func (m *Map) Get(key interface{}) (interface{}, error) {
	value, ok := m.m[key]
	if !ok {
		return value, ErrNotFound
	}
	return value, nil
}

func (m *Map) Remove(key interface{}) {
	if _, ok := m.m[key]; ok {
		delete(m.m, key)
		index, _ := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}

// Container Interface

func (m *Map) Empty() bool {
	return m.Size() == 0
}

func (m *Map) Size() int {
	return m.ordering.Size()
}

func (m *Map) Keys() []interface{} {
	return m.ordering.Values()
}

func (m *Map) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, value := range m.m {
		values = append(values, value)
	}
	return values
}

func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
	m.ordering.Clear()
}
