// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"github.com/prprprus/ds/list/doublelinkedlist"
)

// Iterator represents an iterable structure of the linked hash set.
type Iterator struct {
	s                map[interface{}]int
	internalIterator *doublelinkedlist.Iterator
}

// Iterator returns the iterator object of the linked hash set.
func (s *Set) Iterator() *Iterator {
	return &Iterator{
		s:                s.s,
		internalIterator: s.ordering.Iterator(),
	}
}

// Next returns true if there are still iterable objects, false otherwise.
func (iterator *Iterator) Next() bool {
	return iterator.internalIterator.Next()
}

// Begin reset the iterator to the initial status.
func (iterator *Iterator) Begin() {
	iterator.internalIterator.Begin()
}

// Prev returns true if the prev element exists, false otherwise.
func (iterator *Iterator) Prev() bool {
	return iterator.internalIterator.Prev()
}

// End reset the iterator to the reverse status.
func (iterator *Iterator) End() {
	iterator.internalIterator.End()
}

// Value returns the current value of the linked hash set.
func (iterator *Iterator) Value() interface{} {
	return iterator.internalIterator.Value()
}
