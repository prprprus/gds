// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package arrayqueue implements the array queue.
// Structure is not concurrent safe.
// Reference: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
package arrayqueue

import (
	"errors"

	"github.com/prprprus/ds/list/arraylist"
)

var (
	// ErrPop is returned when the queue is empty
	ErrPop = errors.New("queue is empty")
)

// Queue represents a array queue structure.
type Queue struct {
	list *arraylist.List
}

// New array queue.
func New() *Queue {
	return &Queue{
		list: &arraylist.List{},
	}
}

// Queue Interface

// Put value into the queue.
func (queue *Queue) Put(value interface{}) {
	queue.list.Append(value)
}

// Get the value from the top of the queue and delete the value.
func (queue *Queue) Get() (interface{}, error) {
	if queue.Size() == 0 {
		return nil, ErrPop
	}

	value, _ := queue.list.Get(0)
	queue.list.Remove(0)
	return value, nil
}

// Check if the index is within the length of the list.
func (queue *Queue) indexInRange(index int) bool {
	if index >= 0 && index < queue.list.Size() {
		return true
	}
	return false
}

// Container Interface

// Empty returns true if the queue is empty, otherwise returns false.
func (queue *Queue) Empty() bool {
	return queue.list.Size() == 0
}

// Size returns the size of the queue.
func (queue *Queue) Size() int {
	return queue.list.Size()
}

// Clear the queue.
func (queue *Queue) Clear() {
	queue.list.Clear()
}

// Values returns the values of queue.
func (queue *Queue) Values() []interface{} {
	return queue.list.Values()
}
