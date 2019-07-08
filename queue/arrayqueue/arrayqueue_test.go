// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package arrayqueue

import "testing"

func verifyElements(result []interface{}, queue *Queue) bool {
	values := queue.Values()
	for i, v := range result {
		if values[i] != v {
			return false
		}
	}
	return true
}

func TestNew(t *testing.T) {
	queue := New()
	if queue.Size() != 0 {
		t.Error("TestNew error")
	}
}

// Queue Interface

func TestPut(t *testing.T) {
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	result := []interface{}{1, 2, 3}
	if !verifyElements(result, queue) {
		t.Error("TestPut error")
	}
}

func TestGet(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	v, err := queue.Get()
	if err != nil || v.(int) != 1 {
		t.Error("case1 error: queue has some elements")
	}
	v, err = queue.Get()
	if err != nil || v.(int) != 2 {
		t.Error("case1 error: queue has some elements")
	}
	v, err = queue.Get()
	if err != nil || v.(int) != 3 {
		t.Error("case1 error: queue has some elements")
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	v, err = queue.Get()
	if v.(int) != 1 || err != nil {
		t.Error("case2 error: queue has one element")
	}
	v, err = queue.Get()
	if v != nil || err == nil {
		t.Error("case2 error: queue has one element")
	}

	// case3: queue has no element
	queue = New()
	v, err = queue.Get()
	if v != nil || err == nil {
		t.Error("case3 error: queue has no element")
	}
}

func TestIndexInRange(t *testing.T) {
	// case1: queue has some elements, check with first element index
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	if !queue.indexInRange(0) {
		t.Error("case1 error: queue has some elements, check with first element index")
	}

	// case2: queue has some elements, check with last element index
	queue = New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	if !queue.indexInRange(2) {
		t.Error("case2 error: queue has some elements, check with last element index")
	}

	// case3: queue has some elements, check with medium element index
	queue = New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	if !queue.indexInRange(1) {
		t.Error("case3 error: queue has some elements, check with medium element index")
	}

	// case4: queue has some elements. check with not exists index
	queue = New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	if queue.indexInRange(-1) {
		t.Error("case4 error: queue has some elements. check with not exists index")
	}
	if queue.indexInRange(3) {
		t.Error("case4 error: queue has some elements. check with not exists index")
	}

	// case5: queue has one element
	queue = New()
	queue.Put(1)
	if !queue.indexInRange(0) {
		t.Error("case5 error: queue has one element")
	}
	if queue.indexInRange(1) {
		t.Error("case5 error: queue has one element")
	}

	// case6: queue has no element
	queue = New()
	if queue.indexInRange(0) {
		t.Error("case6 error: queue has no element")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	if queue.Empty() {
		t.Error("case1 error: queue has some elements")
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	if queue.Empty() {
		t.Error("case2 error: queue has one element")
	}

	// case3: queue has no element
	queue = New()
	if !queue.Empty() {
		t.Error("case3 error: queue has no element")
	}
}

func TestSize(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	if queue.Size() != 3 {
		t.Error("case1 error: queue has some elements")
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	if queue.Size() != 1 {
		t.Error("case2 error: queue has one element")
	}

	// case3: queue has no element
	queue = New()
	if queue.Size() != 0 {
		t.Error("case3 error: queue has no element")
	}
}

func TestClear(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	queue.Clear()
	if queue.Size() != 0 {
		t.Error("case1 error: queue has some elements")
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	queue.Clear()
	if queue.Size() != 0 {
		t.Error("case2 error: queue has one element")
	}

	// case3: queue has no element
	queue = New()
	queue.Clear()
	if queue.Size() != 0 {
		t.Error("case3 error: queue has no element")
	}
}

func TestValues(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	result := []interface{}{1, 2, 3}
	if !verifyElements(result, queue) {
		t.Error("case1 error: queue has some elements")
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	result = []interface{}{1}
	if !verifyElements(result, queue) {
		t.Error("case2 error: queue has one element")
	}

	// case3: queue has no element
	queue = New()
	values := queue.Values()
	if len(values) != 0 {
		t.Error("case3 error: queue has no element")
	}
}

// Iterator Interface

func TestIterator(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	iterator := queue.Iterator()
	if iterator.index != -1 {
		t.Error("case1 error: queue has some elements")
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	iterator = queue.Iterator()
	if iterator.index != -1 {
		t.Error("case2 error: queue has one element")
	}

	// case3: queue has no element
	queue = New()
	iterator = queue.Iterator()
	if iterator.index != -1 {
		t.Error("case3 error: queue has no element")
	}
}

func TestNextAndValue(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	iterator := queue.Iterator()
	result := []interface{}{1, 2, 3}
	i := 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case1 error: queue has some elements")
		}
		i++
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	iterator = queue.Iterator()
	result = []interface{}{1}
	i = 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case2 error: queue has one element")
		}
		i++
	}

	// case3: queue has no element
	queue = New()
	iterator = queue.Iterator()
	if iterator.Next() {
		t.Error("case3 error: queue has no element")
	}
}

func TestBegin(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	iterator := queue.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case1 error: queue has some elements")
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	iterator = queue.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case2 error: queue has one element")
	}

	// case3: queue has no element
	queue = New()
	iterator = queue.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case3 error: queue has no element")
	}
}

func TestIndex(t *testing.T) {
	// case1: queue has some elements
	queue := New()
	queue.Put(1)
	queue.Put(2)
	queue.Put(3)
	iterator := queue.Iterator()
	i := 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case1 error: queue has some elements")
		}
		i++
	}

	// case2: queue has one element
	queue = New()
	queue.Put(1)
	iterator = queue.Iterator()
	i = 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case2 error: queue has one element")
		}
		i++
	}

	// case3: queue has no element
	queue = New()
	iterator = queue.Iterator()
	if iterator.index != -1 {
		t.Error("case3 error: queue has no element")
	}
}
