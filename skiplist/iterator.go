// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package skiplist

// Iterator represents an iterable structure of the skip list.
type Iterator struct {
	skiplist *SkipList
	element  *node
	index    int
}

// Iterator returns the iterator object of the skip list.
func (skiplist *SkipList) Iterator() *Iterator {
	return &Iterator{
		skiplist: skiplist,
		element:  nil,
		index:    -1,
	}
}

// Next returns true if there are still iterable objects, false otherwise.
func (iterator *Iterator) Next() bool {
	if !iterator.skiplist.indexInRange(iterator.index + 1) {
		return false
	}

	if iterator.index == -1 {
		iterator.element = iterator.skiplist.head.next[0]
		iterator.index++
		return true
	}

	iterator.element = iterator.element.next[0]
	iterator.index++
	return true
}

// Begin reset the iterator to the initial status.
func (iterator *Iterator) Begin() {
	iterator.element = nil
	iterator.index = -1
}

// Value returns the current value of the element of the iterator.
func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

// Index returns the current index of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}
