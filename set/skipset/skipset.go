// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package skipset implements the skip set.
//
// Compared to hashset the skip set can additionally
// maintain the order of key-value by the skip list
// and compared to linkedhashset the skip set performance is better,
// the time complexity of put, get, remove operations is O(log N).
//
// Structure is not concurrent safe.
// TODO: Add more methods related to the key-value order.
package skipset

import (
	"errors"

	"github.com/prprprus/ds/skiplist"
)

const (
	// FixValue is the value of map, fixed at 1
	FixValue = 1
)

var (
	// ErrEmpty is returned when the skip set is empty
	ErrEmpty = errors.New("the skip set is empty")
)

// The Set represents a skip set structure.
type Set struct {
	s *skiplist.SkipList
}

// New the skip set.
func New(comparator func(a, b interface{}) int) *Set {
	return &Set{
		s: skiplist.New(comparator),
	}
}

// Set Interface

// Add the values into the skip set.
func (s *Set) Add(values ...interface{}) {
	for _, v := range values {
		if !s.s.Exists(v) {
			s.s.Set(v, FixValue)
		}
	}
}

// Remove the values into the skip set.
func (s *Set) Remove(values ...interface{}) error {
	if s.Size() == 0 {
		return ErrEmpty
	}

	for _, v := range values {
		s.s.Remove(v)
	}
	return nil
}

// Contains returns true if the skip set contains values, otherwise returns false.
func (s *Set) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if s.Size() == 0 {
		return false
	}

	for _, v := range values {
		if !s.s.Exists(v) {
			return false
		}
	}
	return true
}

// Container Interface

// Empty returns true if the skip set is empty, otherwise returns false.
func (s *Set) Empty() bool {
	return s.Size() == 0
}

// Size returns the size of the skip set.
func (s *Set) Size() int {
	return s.s.Size()
}

// Clear the skip set.
func (s *Set) Clear() {
	s.s.Clear()
}

// Values returns the values of the skip set.
func (s *Set) Values() []interface{} {
	values := make([]interface{}, 0)
	iterator := s.s.Iterator()
	for iterator.Next() {
		values = append(values, iterator.Key())
	}
	return values
}
