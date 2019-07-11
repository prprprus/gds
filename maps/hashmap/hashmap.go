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
	m map[interface{}]interface{}
}

// New the hash map.
func New() *Map {
	return &Map{
		m: make(map[interface{}]interface{}),
	}
}

// Map Interface

// Put the key-value into map.
func (m *Map) Put(key, value interface{}) {
	m.m[key] = value
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
func (m *Map) Remove(key interface{}) {
	delete(m.m, key)
}

// Container Interface

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, 0)
	for key := range m.m {
		keys = append(keys, key)
	}
	return keys
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, 0)
	for _, value := range m.m {
		values = append(values, value)
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
}
