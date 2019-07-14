// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package hashmap implements the hash map.
// Structure is not concurrent safe.
package hashmap

import "errors"

var (
	// ErrNotFound is returned when the value can not found
	ErrNotFound = errors.New("can not found value")
	// ErrEmpty is returned when the hash map is empty
	ErrEmpty = errors.New("the hash map is empty")
)

// The Map represents a hash map structure.
type Map struct {
	m    map[interface{}]interface{}
	size int
}

// New the hash map.
func New() *Map {
	return &Map{
		m: make(map[interface{}]interface{}),
	}
}

// Map Interface

// Put the key-value into the hash map.
func (m *Map) Put(key, value interface{}) {
	m.m[key] = value
	m.size++
}

// Get the value by key.
func (m *Map) Get(key interface{}) (interface{}, error) {
	value, ok := m.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return value, nil
}

// Remove key-value by key.
func (m *Map) Remove(key interface{}) error {
	if _, ok := m.m[key]; !ok {
		return ErrNotFound
	}
	delete(m.m, key)
	m.size--
	return nil
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, 0)
	for key := range m.m {
		keys = append(keys, key)
	}
	return keys
}

// Container Interface

// Empty returns true if hash map does not contain any elements.
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the hash map.
func (m *Map) Size() int {
	return len(m.m)
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, value := range m.m {
		values = append(values, value)
	}
	return values
}

// Clear removes all elements from the hash map.
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
	m.size = 0
}
