// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package arraystack

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

// Prev returns true if the prev element exists, false otherwise.
func (iterator *Iterator) Prev() bool {
	if !iterator.stack.indexInRange(iterator.index - 1) {
		return false
	}

	iterator.index--
	return true
}

// End reset the iterator to the reverse status.
func (iterator *Iterator) End() {
	iterator.index = iterator.stack.Size()
}

// Value returns the current value of the element of the iterator.
func (iterator *Iterator) Value() interface{} {
	if !iterator.stack.indexInRange(iterator.index) {
		return nil
	}

	value, _ := iterator.stack.list.Get(iterator.stack.Size() - iterator.index - 1)
	return value
}

// Index returns the current index of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}
