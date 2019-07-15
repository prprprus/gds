// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package list

import (
	"github.com/prprprus/ds/container"
	"github.com/prprprus/ds/iterator"
)

// List interface
type List interface {
	Append(values ...interface{})
	Get(index int) (interface{}, error)
	Remove(index int) error
	Contains(values ...interface{}) bool
	Swap(i, j int) error
	Insert(index int, values ...interface{}) error
	Set(index int, value interface{}) error
	IndexOf(value interface{}) (int, error)

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}

	iterator.IndexIterator
	// ValueIterator
	// 	Next() bool
	// 	Begin()
	// 	Value() interface{}
	// Index() int
}
