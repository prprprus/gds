// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package singlylinkedlist implements the singly-linked list.
// Structure is not concurrent safe.
// Reference: https://en.wikipedia.org/wiki/Linked_list#Singly_linked_list
package singlylinkedlist

import (
	"errors"
)

var (
	// ErrIndex is returned when the index is out of the list
	ErrIndex = errors.New("index is out of the list")
	// ErrIndexOf is returned when the index of value can not found
	ErrIndexOf = errors.New("index of value can not found")
)

// List represents a singly linked list structure.
type List struct {
	first *element // the address of the first element
	last  *element // the address of the last element
	size  int      // size of list
}

// element of list.
type element struct {
	value interface{} // value fields can store any type
	next  *element    // next fields stores the address pointing to the next element
}

// New singly linked list.
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) != 0 {
		list.Append(values...)
	}
	return list
}

// List Interface

// Append values (one or more than one) to list.
func (list *List) Append(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	// if the size is equal to 0, mean it is a new singly linked list
	if list.size == 0 {
		for i, v := range values {
			newElement := &element{value: v}
			// first element
			if i == 0 {
				list.first = newElement
				list.last = newElement
			} else {
				list.last.next = newElement
				list.last = newElement
			}
			list.size++
		}
	} else {
		for _, v := range values {
			newElement := &element{value: v}
			list.last.next = newElement
			list.last = newElement
			list.size++
		}
	}
}

// PreAppend can append values (one or more than one) to the front of the list.
func (list *List) PreAppend(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	if list.size == 0 {
		list.Append(values...)
	} else {
		for i := len(values) - 1; i >= 0; i-- {
			newElement := &element{value: values[i]}
			newElement.next = list.first
			list.first = newElement
			list.size++
		}
	}
}

// indexInRange check if the index is within the length of the list.
func (list *List) indexInRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}

// Get value by index.
func (list *List) Get(index int) (interface{}, error) {
	if !list.indexInRange(index) {
		return nil, ErrIndex
	}

	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}

	return foundElement.value, nil
}

// Remove element by index.
func (list *List) Remove(index int) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	foundElement := list.first

	if index == 0 {
		list.first = list.first.next
	} else {
		preFoundElement := new(element)
		for i := 0; i != index; i++ {
			preFoundElement = foundElement
			foundElement = foundElement.next
		}
		preFoundElement.next = foundElement.next
	}

	foundElement.value = nil
	foundElement.next = nil
	list.size--

	return nil
}

// Contains returns true if list contains values, false otherwise.
func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}

	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Swap value by index.
func (list *List) Swap(i, j int) error {
	if !list.indexInRange(i) || !list.indexInRange(j) {
		return ErrIndex
	}
	if i == j {
		return nil
	}

	foundElementI := list.first
	for index := 0; index != i; index++ {
		foundElementI = foundElementI.next
	}

	foundElementJ := list.first
	for index := 0; index != j; index++ {
		foundElementJ = foundElementJ.next
	}

	foundElementI.value, foundElementJ.value = foundElementJ.value, foundElementI.value

	return nil
}

// Insert value (one or more than one) after index.
func (list *List) Insert(index int, values ...interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}
	if len(values) == 0 {
		return nil
	}
	if index == list.size-1 {
		list.Append(values...)
		return nil
	}

	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElementNext := foundElement.next

	// insert
	for _, v := range values {
		element := &element{value: v}
		foundElement.next = element
		foundElement = element
	}
	foundElement.next = foundElementNext
	list.size += len(values)

	return nil
}

// Set element by index.
func (list *List) Set(index int, value interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	foundElement := list.first
	for i := 0; i != index; i++ {
		foundElement = foundElement.next
	}
	foundElement.value = value

	return nil
}

// IndexOf get index by value.
func (list *List) IndexOf(value interface{}) (int, error) {
	for i, v := range list.Values() {
		if v == value {
			return i, nil
		}
	}
	return -1, ErrIndexOf
}

// Reverse the list.
func (list *List) Reverse() {
	if list.size == 0 || list.size == 1 {
		return
	}

	preElement := new(element)
	preElement.value = nil
	preElement.next = nil
	currElement := list.first
	nextElement := list.first.next

	// reset the last pointer
	list.last = currElement

	// reverse
	for currElement != nil {
		currElement.next = preElement
		preElement = currElement
		currElement = nextElement
		if nextElement != nil {
			nextElement = nextElement.next
		}
	}

	// reset the first pointer
	list.first = preElement
}

// Container Interface

// Empty returns true if the list is empty, otherwise returns false.
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns the size of the list.
func (list *List) Size() int {
	size := list.size
	return size
}

// Clear th list.
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Values returns the values of list.
func (list *List) Values() []interface{} {
	values := make([]interface{}, 0)

	if list.size == 0 {
		return values
	}

	iterator := list.Iterator()
	iterator.Begin()
	for iterator.Next() {
		values = append(values, iterator.Value())
	}

	return values
}
