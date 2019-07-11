// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package skipmap

import "github.com/prprprus/ds/skiplist"

// Iterator represents an iterable structure of the skip map.
type Iterator struct {
	m                *skiplist.SkipList
	internalIterator *skiplist.Iterator
}

// Iterator returns the iterator object of the skip map.
func (m *Map) Iterator() *Iterator {
	return &Iterator{
		m:                m.m,
		internalIterator: m.m.Iterator(),
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

// Key returns the current key of the skip map.
func (iterator *Iterator) Key() interface{} {
	return iterator.internalIterator.Key()
}

// Value returns the current value of the skip map.
func (iterator *Iterator) Value() interface{} {
	return iterator.internalIterator.Value()
}
