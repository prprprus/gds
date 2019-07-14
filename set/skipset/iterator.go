// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package skipset

import "github.com/prprprus/ds/skiplist"

// Iterator represents an iterable structure of the skip set.
type Iterator struct {
	s                *skiplist.SkipList
	internalIterator *skiplist.Iterator
}

// Iterator returns the iterator object of the skip set.
func (s *Set) Iterator() *Iterator {
	return &Iterator{
		s:                s.s,
		internalIterator: s.s.Iterator(),
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

// Value returns the current value of the skip set.
func (iterator *Iterator) Value() interface{} {
	return iterator.internalIterator.Key()
}
