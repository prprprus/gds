// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package hashmap

import "testing"

func TestNew(t *testing.T) {
	m := New()
	if m.m == nil {
		t.Error("TestNew error")
	}
}

// Map Interface

func TestPut(t *testing.T) {
	// case1: map has some elements
	m := New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if len(m.m) != 3 || m.m[1] != 1 || m.m[2] != 2 || m.m[3] != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New()
	m.Put(1, 1)
	if len(m.m) != 1 || m.m[1] != 1 {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New()
	if len(m.m) != 0 {
		t.Error("case3 error: map has one element")
	}
}

func TestGet(t *testing.T) {
	// case1: map has some elements
	m := New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	v1, _ := m.Get(1)
	v2, _ := m.Get(2)
	v3, _ := m.Get(3)
	if len(m.m) != 3 || v1.(int) != 1 || v2.(int) != 2 || v3.(int) != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New()
	m.Put(1, 1)
	v1, _ = m.Get(1)
	if len(m.m) != 1 || v1.(int) != 1 {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New()
	_, err := m.Get(1)
	if len(m.m) != 0 || err == nil {
		t.Error("case3 error: map has one element")
	}
}

func TestRemove(t *testing.T) {
	// case1: map has some elements
	m := New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Remove(2)
	m.Remove(1)
	_, err1 := m.Get(1)
	_, err2 := m.Get(2)
	if len(m.m) != 1 || err1 == nil || err2 == nil || m.m[3] != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New()
	m.Put(1, 1)
	m.Remove(1)
	_, err := m.Get(1)
	if len(m.m) != 0 || err == nil {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New()
	m.Remove(1)
	if len(m.m) != 0 {
		t.Error("case3 error: map has one element")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: the map has no element
	m := New()
	if !m.Empty() {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New()
	m.Put(1, 1)
	if m.Empty() {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if m.Empty() {
		t.Error("case3 error: the map has some elements")
	}
}

func TestSize(t *testing.T) {
	// case1: the map has no element
	m := New()
	if m.Size() != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New()
	m.Put(1, 1)
	if m.Size() != 1 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if m.Size() != 3 {
		t.Error("case3 error: the map has some elements")
	}
}

func TestClear(t *testing.T) {
	// case1: the map has no element
	m := New()
	m.Clear()
	if m.size != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New()
	m.Put(1, 1)
	m.Clear()
	if m.size != 0 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Clear()
	if m.size != 0 {
		t.Error("case3 error: the map has some elements")
	}
}

func TestValues(t *testing.T) {
	// case1: the map has no element
	m := New()
	values := m.Values()
	if len(values) != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	values = m.Values()
	if len(values) != 1 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	values = m.Values()
	if len(values) != 3 {
		t.Error("case3 error: the map has some elements")
	}
}
