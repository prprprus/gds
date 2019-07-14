// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package skipmap implements the skip map.
//
// Compared to hashmap the skip map can additionally
// maintain the order of key-value by the skip list
// and compared to linkedhashmap the skip map performance is better,
// the time complexity of put, get, remove operations is O(log N).
//
// Structure is not concurrent safe.
// TODO: Add more methods related to the key-value order.
package skipmap

import (
	"errors"

	"github.com/prprprus/ds/skiplist"
)

var (
	// ErrNotFound is returned when the value can not found
	ErrNotFound = errors.New("can not found value")
	// ErrEmpty is returned when the skip map is empty
	ErrEmpty = errors.New("the skip map is empty")
)

// The Map represents a skip map structure.
type Map struct {
	m *skiplist.SkipList
}

// New the skip map.
func New(comparator func(a, b interface{}) int) *Map {
	return &Map{
		m: skiplist.New(comparator),
	}
}

// Put the key-value into the skip map.
func (m *Map) Put(key, value interface{}) {
	m.m.Set(key, value)
}

// Get the value by key.
func (m *Map) Get(key interface{}) (interface{}, error) {
	return m.m.Get(key)
}

// Remove key-value by key.
func (m *Map) Remove(key interface{}) {
	m.m.Remove(key)
}

// Keys returns all keys (orderly).
func (m *Map) Keys() []interface{} {
	return m.m.Values()
}

// Container Interface

// Empty returns true if skip map does not contain any elements.
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the skip map.
func (m *Map) Size() int {
	return m.m.Size()
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

// Clear removes all elements from the skip map.
func (m *Map) Clear() {
	m.m.Clear()
}
