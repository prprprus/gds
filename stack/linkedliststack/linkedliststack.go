// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package linkedliststack implements the linked list stack.
// Structure is not concurrent safe.
// Reference: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)#Linked_list
package linkedliststack

import (
	"errors"

	"github.com/prprprus/ds/list/singlylinkedlist"
)

var (
	// ErrPop is returned when the stack is empty
	ErrPop = errors.New("stack is empty")
)

// Stack represents a linked list stack structure.
type Stack struct {
	list *singlylinkedlist.List
}

// New linked list stack.
func New() *Stack {
	return &Stack{
		list: &singlylinkedlist.List{},
	}
}

// Stack Interface

// Push value into the stack.
func (stack *Stack) Push(value interface{}) {
	stack.list.PreAppend(value)
}

// Pop the value from the top of the stack and delete the value.
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Size() == 0 {
		return nil, ErrPop
	}

	value, _ := stack.list.Get(0)
	stack.list.Remove(0)
	return value, nil
}

// Peek will get the value from the top of the stack and don't delete the value.
func (stack *Stack) Peek() (interface{}, error) {
	if stack.Size() == 0 {
		return nil, ErrPop
	}

	value, _ := stack.list.Get(0)
	return value, nil
}

// Check if the index is within the length of the list.
func (stack *Stack) indexInRange(index int) bool {
	if index >= 0 && index < stack.list.Size() {
		return true
	}
	return false
}

// Container Interface

// Empty returns true if the stack is empty, otherwise returns false.
func (stack *Stack) Empty() bool {
	return stack.list.Size() == 0
}

// Size returns the size of the stack.
func (stack *Stack) Size() int {
	return stack.list.Size()
}

// Clear the stack.
func (stack *Stack) Clear() {
	stack.list.Clear()
}

// Values returns the values of stack.
func (stack *Stack) Values() []interface{} {
	return stack.list.Values()
}
