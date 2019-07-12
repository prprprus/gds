// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package linkedhashmap implements the linked hash map.
//
// Compared to hashmap the linked hash map can additionally
// maintain the order of key-value by the double linked list.
//
// Structure is not concurrent safe.
// TODO: Add more methods related to the key-value order.
package linkedhashmap

import (
	"errors"

	"github.com/prprprus/ds/list/doublelinkedlist"
)

var (
	// ErrNotFound is returned when the value can not found
	ErrNotFound = errors.New("can not found value")
	// ErrEmpty is returned when the linked hash map is empty
	ErrEmpty = errors.New("the linked hash map is empty")
)

// The Map represents a linked hash map structure.
type Map struct {
	m        map[interface{}]interface{}
	ordering *doublelinkedlist.List // maintain the order of key-value
}

// New the linked hash map.
func New() *Map {
	return &Map{
		m:        make(map[interface{}]interface{}),
		ordering: doublelinkedlist.New(),
	}
}

// Map Interface

// Put the key-value into linked hash map.
func (m *Map) Put(key, value interface{}) {
	if _, ok := m.m[key]; !ok {
		m.ordering.Append(value)
	}
	m.m[key] = value
}

// Get the value by key.
func (m *Map) Get(key interface{}) (interface{}, error) {
	value, ok := m.m[key]
	if !ok {
		return value, ErrNotFound
	}
	return value, nil
}

// Remove key-value by key.
func (m *Map) Remove(key interface{}) {
	if _, ok := m.m[key]; ok {
		delete(m.m, key)
		index, _ := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}

// Container Interface

// Empty returns true if linked hash map does not contain any elements.
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the linked hash map.
func (m *Map) Size() int {
	return m.ordering.Size()
}

// Keys returns all keys (orderly).
func (m *Map) Keys() []interface{} {
	return m.ordering.Values()
}

// Values returns all values (orderly).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	iterator := m.Iterator()
	for iterator.Next() {
		values[count] = iterator.Value()
		count++
	}
	return values
}

// Clear removes all elements from the linked hash map.
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
	m.ordering.Clear()
}
