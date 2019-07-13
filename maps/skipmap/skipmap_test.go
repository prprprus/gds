// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package skipmap

import (
	"testing"

	"github.com/prprprus/ds/util"
)

func TestNew(t *testing.T) {
	m := New(util.IntComparator)
	if m.m == nil {
		t.Error("TestNew error")
	}
}

// Map Interface

func TestPut(t *testing.T) {
	// case1: map has some elements
	m := New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	v1, _ := m.m.Get(1)
	v2, _ := m.m.Get(2)
	v3, _ := m.m.Get(3)
	if m.m.Size() != 3 || v1.(int) != 1 || v2.(int) != 2 || v3.(int) != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	v1, _ = m.m.Get(1)
	if m.m.Size() != 1 || v1.(int) != 1 {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New(util.IntComparator)
	if m.m.Size() != 0 {
		t.Error("case3 error: map has one element")
	}
}

func TestGet(t *testing.T) {
	// case1: map has some elements
	m := New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	v1, _ := m.Get(1)
	v2, _ := m.Get(2)
	v3, _ := m.Get(3)
	if m.m.Size() != 3 || v1.(int) != 1 || v2.(int) != 2 || v3.(int) != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	v1, _ = m.Get(1)
	if m.m.Size() != 1 || v1.(int) != 1 {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New(util.IntComparator)
	_, err := m.Get(1)
	if m.m.Size() != 0 || err == nil {
		t.Error("case3 error: map has one element")
	}
}

func TestRemove(t *testing.T) {
	// case1: map has some elements
	m := New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Remove(2)
	m.Remove(1)
	_, err1 := m.Get(1)
	_, err2 := m.Get(2)
	v, _ := m.Get(3)
	if m.m.Size() != 1 || err1 == nil || err2 == nil || v.(int) != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Remove(1)
	_, err := m.Get(1)
	if m.m.Size() != 0 || err == nil {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New(util.IntComparator)
	m.Remove(1)
	if m.m.Size() != 0 {
		t.Error("case3 error: map has one element")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	if !m.Empty() {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	if m.Empty() {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if m.Empty() {
		t.Error("case3 error: the map has some elements")
	}
}

func TestSize(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	if m.Size() != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	if m.Size() != 1 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if m.Size() != 3 {
		t.Error("case3 error: the map has some elements")
	}
}

func TestClear(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	m.Clear()
	if m.m.Size() != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Clear()
	if m.m.Size() != 0 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Clear()
	if m.m.Size() != 0 {
		t.Error("case3 error: the map has some elements")
	}
}

func TestValues(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	values := m.Values()
	if len(values) != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	values = m.Values()
	if len(values) != 1 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	values = m.Values()
	if len(values) != 3 {
		t.Error("case3 error: the map has some elements")
	}
}

// Iterator Interface

func TestIterator(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	iterator := m.Iterator()
	if iterator.m == nil || iterator.internalIterator == nil {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	iterator = m.Iterator()
	if iterator.m == nil || iterator.internalIterator == nil {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	if iterator.m == nil || iterator.internalIterator == nil {
		t.Error("case3 error: the map has some elements")
	}
}

func TestNext(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	iterator := m.Iterator()
	if iterator.Next() {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	iterator = m.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case2 error: the map has one element")
		}
		i++
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	result = []interface{}{1, 2, 3, 4}
	i = 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case3 error: the map has some elements")
		}
		i++
	}
}

func TestBegin(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	iterator := m.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.internalIterator.Index() != -1 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	iterator = m.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.internalIterator.Index() != -1 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.internalIterator.Index() != -1 {
		t.Error("case3 error: the map has some elements")
	}
}

func TestValue(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	iterator := m.Iterator()
	iterator.Next()
	if iterator.Value() != nil {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	iterator = m.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case2 error: the map has one element")
		}
		i++
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	iterator = m.Iterator()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	result = []interface{}{1, 2, 3, 4}
	i = 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case3 error: the map has some elements")
		}
		i++
	}
}

func TestKey(t *testing.T) {
	// case1: the map has no element
	m := New(util.IntComparator)
	iterator := m.Iterator()
	if iterator.Key() != nil {
		t.Error("case1 error: the m has no element")
	}

	// case2: the map has one element
	m = New(util.IntComparator)
	m.Put(1, 1)
	iterator = m.Iterator()
	i := 1
	for iterator.Next() {
		if iterator.Key() != i {
			t.Error("case2 error: the map has one element")
		}
		i++
	}

	// case3: the map has some elements
	m = New(util.IntComparator)
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	i = 1
	for iterator.Next() {
		if iterator.Key() != i {
			t.Error("case3 error: the map has some elements")
		}
		i++
	}
}

// Benchmark Test

// Put

func benchmarkPut(b *testing.B, m *Map, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			m.Put(i, i)
		}
	}
}

func BenchmarkGetPut0(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 0
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkPut100(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 100
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkPut1000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 1000
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkPut10000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 10000
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkPut100000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 100000
	b.StartTimer()
	benchmarkPut(b, m, size)
}

// Get

func benchmarkGet(b *testing.B, m *Map, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			m.Get(i)
		}
	}
}

func BenchmarkGetGet0(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 0
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkGet100(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 100
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkGet1000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 1000
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkGet10000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 10000
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkGet100000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 100000
	b.StartTimer()
	benchmarkGet(b, m, size)
}

// Remove

func benchmarkRemove(b *testing.B, m *Map, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			m.Remove(i)
		}
	}
}

func BenchmarkRemoveRemove0(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 0
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkRemove100(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 100
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkRemove1000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 1000
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkRemove10000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 10000
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkRemove100000(b *testing.B) {
	b.StopTimer()
	m := New(util.IntComparator)
	size := 100000
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
