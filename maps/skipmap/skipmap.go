package skipmap

import (
	"errors"

	"github.com/prprprus/ds/skiplist"
)

var (
	// ErrNotFound is returned when the value can not found
	ErrNotFound = errors.New("can not found value")
	// ErrEmpty is returned when the map is empty
	ErrEmpty = errors.New("the map is empty")
)

type Map struct {
	m *skiplist.SkipList
}

func New(comparator func(a, b interface{}) int) *Map {
	return &Map{
		m: skiplist.New(comparator),
	}
}

func (m *Map) Put(key, value interface{}) {
	m.m.Set(key, value)
}

func (m *Map) Get(key interface{}) (interface{}, error) {
	return m.m.Get(key)
}

func (m *Map) Remove(key interface{}) {
	m.m.Remove(key)
}

// Container Interface

// Empty returns true if map does not contain any elements.
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return m.m.Size()
}

// Keys returns all keys (orderly).
func (m *Map) Keys() []interface{} {
	return m.m.Values()
}

// Values returns all values (orderly).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		values[count] = it.Value()
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.m.Clear()
}
