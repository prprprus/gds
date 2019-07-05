// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package arraylist

// Iterator represents an iterable structure of list.
type Iterator struct {
	list  *List // target of iteration
	index int   // current index
}

// Iterator returns the iterator object of list.
func (list *List) Iterator() *Iterator {
	return &Iterator{
		list:  list,
		index: -1,
	}
}

// Begin reset the iterator to the initial status.
func (iterator *Iterator) Begin() {
	iterator.index = -1
}

// Next returns true if the next element exists, false otherwise.
func (iterator *Iterator) Next() bool {
	if !iterator.list.indexInRange(iterator.index + 1) {
		return false
	}

	iterator.index++
	return true
}

// End reset the iterator to the reverse status.
func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
}

// Prev returns true if the prev element exists, false otherwise.
func (iterator *Iterator) Prev() bool {
	if !iterator.list.indexInRange(iterator.index - 1) {
		return false
	}

	iterator.index--
	return true
}

// Value returns the current value of the element of the iterator.
func (iterator *Iterator) Value() interface{} {
	if !iterator.list.indexInRange(iterator.index) {
		return nil
	}

	return iterator.list.elements[iterator.index]
}

// Index returns the current index of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}
