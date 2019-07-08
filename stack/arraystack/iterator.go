// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package arraystack

type Iterator struct {
	stack *Stack
	index int
}

func (stack *Stack) Iterator() *Iterator {
	return &Iterator{
		stack: stack,
		index: -1,
	}
}

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

func (iterator *Iterator) Prev() bool {
	if !iterator.stack.indexInRange(iterator.index - 1) {
		return false
	}

	iterator.index--
	return true
}

func (iterator *Iterator) End() {
	iterator.index = iterator.stack.Size()
}

func (iterator *Iterator) Value() interface{} {
	if !iterator.stack.indexInRange(iterator.index) {
		return nil
	}

	value, _ := iterator.stack.list.Get(iterator.stack.Size() - iterator.index - 1)
	return value
}

func (iterator *Iterator) Index() int {
	return iterator.index
}
