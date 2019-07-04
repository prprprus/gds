// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package doublelinkedlist

// Iterator represents an iterable structure of list.
type Iterator struct {
	list    *List    // target of iteration
	index   int      // current index
	element *element // current element
}

// Iterator returns the iterator object of list.
func (list *List) Iterator() *Iterator {
	return &Iterator{
		list:    list,
		index:   -1,
		element: nil,
	}
}

// Begin reset the iterator to the initial status.
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// Next returns true if the next element exists, false otherwise.
func (iterator *Iterator) Next() bool {
	if !iterator.list.indexInRange(iterator.index + 1) {
		return false
	}

	if iterator.index == -1 {
		iterator.index++
		iterator.element = iterator.list.first
		return true
	}

	iterator.index++
	iterator.element = iterator.element.next
	return true
}

// End reset the iterator to the reverse status.
func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
	iterator.element = iterator.list.last
}

// Prev returns true if the prev element exists, false otherwise.
func (iterator *Iterator) Prev() bool {
	if !iterator.list.indexInRange(iterator.index - 1) {
		return false
	}

	if iterator.index == iterator.list.size {
		iterator.index--
		iterator.element = iterator.list.last
		return true
	}

	iterator.index--
	iterator.element = iterator.element.prev
	return true
}

// Value returns the current value of the element of the iterator.
func (iterator *Iterator) Value() interface{} {
	if iterator.element == nil {
		return nil
	}

	return iterator.element.value
}

// Index returns the current index of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}
