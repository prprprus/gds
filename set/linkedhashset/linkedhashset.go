// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package linkedhashset implements the linked hash set.
//
// Compared to hashset the linked hash set can additionally
// maintain the order of key-value by the double linked list.
//
// Structure is not concurrent safe.
package linkedhashset

import (
	"errors"

	"github.com/prprprus/ds/list/doublelinkedlist"
)

const (
	// FixValue is the value of map, fixed at 1
	FixValue = 1
)

var (
	// ErrEmpty is returned when the linked hash set is empty
	ErrEmpty = errors.New("the hash set is empty")
)

// The Set represents a linked hash set structure.
type Set struct {
	s        map[interface{}]int
	ordering *doublelinkedlist.List
	size     int
}

// New the linked hash set.
func New() *Set {
	return &Set{
		s:        make(map[interface{}]int),
		ordering: doublelinkedlist.New(),
		size:     0,
	}
}

// Set Interface

// Add the values into the linked hash set.
func (s *Set) Add(values ...interface{}) {
	for _, v := range values {
		if _, ok := s.s[v]; !ok {
			s.s[v] = FixValue
			s.ordering.Append(v)
			s.size++
		}
	}
}

// Remove the values into the linked hash set.
func (s *Set) Remove(values ...interface{}) error {
	if s.Size() == 0 {
		return ErrEmpty
	}

	for _, v := range values {
		delete(s.s, v)
		index, _ := s.ordering.IndexOf(v)
		s.ordering.Remove(index)
		s.size--
	}
	return nil
}

// Contains returns true if the linked hash set contains values, otherwise returns false.
func (s *Set) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if s.Size() == 0 {
		return false
	}

	for _, v := range values {
		if _, ok := s.s[v]; !ok {
			return false
		}
	}
	return true
}

// Container Interface

// Empty returns true if the linked hash set is empty, otherwise returns false.
func (s *Set) Empty() bool {
	return s.Size() == 0
}

// Size returns the size of the linked hash set.
func (s *Set) Size() int {
	return s.size
}

// Clear the linked hash set.
func (s *Set) Clear() {
	s.s = make(map[interface{}]int)
	s.ordering.Clear()
	s.size = 0
}

// Values returns the values of the linked hash set.
func (s *Set) Values() []interface{} {
	return s.ordering.Values()
}
