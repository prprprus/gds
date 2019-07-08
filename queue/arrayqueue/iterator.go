// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package arrayqueue

// Iterator represents an iterable structure of queue.
type Iterator struct {
	queue *Queue
	index int
}

// Iterator returns the iterator object of queue.
func (queue *Queue) Iterator() *Iterator {
	return &Iterator{
		queue: queue,
		index: -1,
	}
}

// Next returns true if the next element exists, false otherwise.
func (iterator *Iterator) Next() bool {
	if !iterator.queue.indexInRange(iterator.index + 1) {
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
	if !iterator.queue.indexInRange(iterator.index) {
		return nil
	}

	value, _ := iterator.queue.list.Get(iterator.index)
	return value
}

// Index returns the current index of the iterator.
func (iterator *Iterator) Index() int {
	return iterator.index
}
