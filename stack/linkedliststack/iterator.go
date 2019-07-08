// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package linkedliststack

// Iterator represents an iterable structure of stack.
type Iterator struct {
	stack *Stack
	index int
}

// Iterator returns the iterator object of stack.
func (stack *Stack) Iterator() *Iterator {
	return &Iterator{
		stack: stack,
		index: -1,
	}
}

// Next returns true if the next element exists, false otherwise.
func (iterator *Iterator) Next() bool {
	if !iterator.stack.indexInRange(iterator.index + 1) {
		return false
	}

	iterator.index++
	return true
}

// Begin reset the iterator to the initial status.
func (iterator *Iterator) Begin() {
	iterator.index = -1
}

// Value returns the current value of the element of the iterator.
func (iterator *Iterator) Value() interface{} {
	value, _ := iterator.stack.list.Get(iterator.index)
	return value
}

// Index returns the current index of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}
