// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package hashset implements the hash set.
// Structure is not concurrent safe.
package hashset

import (
	"errors"
)

const (
	// FixValue is the value of map, fixed at 1
	FixValue = 1
)

var (
	// ErrEmpty is returned when the hash set is empty
	ErrEmpty = errors.New("the hash set is empty")
)

// The Set represents a hash set structure.
type Set struct {
	s    map[interface{}]int
	size int
}

// New the hash set.
func New() *Set {
	return &Set{
		s:    make(map[interface{}]int),
		size: 0,
	}
}

// Set Interface

// Add the values into the hash set.
func (s *Set) Add(values ...interface{}) {
	for _, v := range values {
		s.s[v] = FixValue
		s.size++
	}
}

// Remove the values into the hash set.
func (s *Set) Remove(values ...interface{}) error {
	if s.Size() == 0 {
		return ErrEmpty
	}

	for _, v := range values {
		delete(s.s, v)
		s.size--
	}
	return nil
}

// Contains returns true if the hash set contains values, otherwise returns false.
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

// Empty returns true if the hash set is empty, otherwise returns false.
func (s *Set) Empty() bool {
	return s.Size() == 0
}

// Size returns the size of the hash set.
func (s *Set) Size() int {
	return s.size
}

// Clear the hash set.
func (s *Set) Clear() {
	s.s = make(map[interface{}]int)
	s.size = 0
}

// Values returns the values of the hash set.
func (s *Set) Values() []interface{} {
	values := make([]interface{}, 0)
	for v := range s.s {
		values = append(values, v)
	}
	return values
}
