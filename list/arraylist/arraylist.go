// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package arraylist implements the array list.
// Structure is not concurrent safe.
// Reference: https://en.wikipedia.org/wiki/Dynamic_array
package arraylist

import (
	"errors"
)

var (
	// ErrIndex is returned when the index is out of the list
	ErrIndex = errors.New("index is out of the list")
	// ErrIndexOf is returned when the index of value can not found
	ErrIndexOf = errors.New("index of value can not found")
)

const (
	growthFactor = float32(2.0)  // growth factor by 100%
	shrinkFactor = float32(0.25) // shrink factor by 25%
)

// List represents a array list structure.
type List struct {
	elements []interface{} // elements of array list
	size     int           // size of array list
	caps     int           // capacity of array list
}

// Growth the list capacity.
// If the number of elements is greater than the current capacity,
// expand the capacity by 100%.
func (list *List) growth(n int) {
	currCap := cap(list.elements)
	if list.size+n >= currCap {
		newCap := int(float32(currCap+n) * growthFactor)
		list.resize(newCap)
	}
}

// Shrink the list capacity.
// If the number of elements is less than 25% of the current capacity,
// reduce the capacity to number of elements.
func (list *List) shrink() {
	// shrinkFactor equal 0.0 mean never shrink
	if shrinkFactor == 0.0 {
		return
	}

	currCap := cap(list.elements)
	if list.size <= int(float32(currCap)*shrinkFactor) {
		list.resize(list.size)
	}
}

// Resize the list capacity.
func (list *List) resize(caps int) {
	newElements := make([]interface{}, caps, caps)
	copy(newElements, list.elements)
	list.elements = newElements
	list.caps = caps
}

// New array list.
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
	list.growth(len(values))
	for _, v := range values {
		list.elements[list.size] = v
		list.size++
	}
}

// Check if the index is within the length of the list.
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

	return list.elements[index], nil
}

// Remove element by index.
func (list *List) Remove(index int) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	// remove
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()

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
		for _, element := range list.elements {
			if element == value {
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

	list.elements[i], list.elements[j] = list.elements[j], list.elements[i]

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

	// insert
	l := len(values)
	list.growth(l)
	list.size += l
	copy(list.elements[index+1+l:], list.elements[index+1:])
	copy(list.elements[index+1:], values)

	return nil
}

// Set element by index.
func (list *List) Set(index int, value interface{}) error {
	if !list.indexInRange(index) {
		return ErrIndex
	}

	list.elements[index] = value

	return nil
}

// IndexOf get index by value.
func (list *List) IndexOf(value interface{}) (int, error) {
	for index, element := range list.elements {
		if element == value {
			return index, nil
		}
	}
	return -1, ErrIndexOf
}

// Container Interface

// Empty returns true if the list is empty, otherwise returns false.
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns the size of the list.
func (list *List) Size() int {
	return list.size
}

// Clear the list.
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// Values returns the values of list.
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements)
	return newElements
}
