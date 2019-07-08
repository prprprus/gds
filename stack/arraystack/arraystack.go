// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package arraystack implements the array stack.
// Structure is not concurrent safe.
// Reference: https://en.wikipedia.org/wiki/Stack_(abstract_data_type)#Array
package arraystack

import (
	"errors"

	"github.com/prprprus/ds/list/arraylist"
)

var (
	// ErrPop is returned when the stack is empty
	ErrPop = errors.New("stack is empty")
)

// Stack represents a array stack structure.
type Stack struct {
	list *arraylist.List // internal array list
}

// New array stack.
func New() *Stack {
	return &Stack{
		list: arraylist.New(),
	}
}

// Stack Interface

// Push value into the stack.
func (stack *Stack) Push(value interface{}) {
	stack.list.Append(value)
}

// Pop the value from the top of the stack and delete the value.
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Size() == 0 {
		return nil, ErrPop
	}

	element, _ := stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return element, nil
}

// Peek will get the value from the top of the stack and don't delete the value.
func (stack *Stack) Peek() (interface{}, error) {
	if stack.Size() == 0 {
		return nil, ErrPop
	}

	element, _ := stack.list.Get(stack.list.Size() - 1)
	return element, nil
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
