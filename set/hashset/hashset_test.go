// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package hashset

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if s.s == nil {
		t.Error("TestNew error")
	}
}

// Set Interface

func TestAddAndContains(t *testing.T) {
	// case1: set has some elements
	s := New()
	s.Add(1, 2, 3, 4)
	if s.Size() != 4 || !s.Contains(1) || !s.Contains(2) || !s.Contains(3) || !s.Contains(4) {
		t.Error("case1 error: set has some elements")
	}

	// case2: set has one element
	s = New()
	s.Add(1)
	if s.Size() != 1 || !s.Contains(1) {
		t.Error("case2 error: set has one element")
	}

	// case3: set has no element
	s = New()
	if s.Size() != 0 || s.Contains(1) {
		t.Error("case3 error: set has no element")
	}
}

func TestRemoveAndContains(t *testing.T) {
	// case1: set has some elements
	s := New()
	s.Add(1, 2, 3, 4)
	s.Remove(1)
	s.Remove(2)
	if s.Size() != 2 || s.Contains(1) || s.Contains(2) || !s.Contains(3) || !s.Contains(4) {
		t.Error("case1 error: set has some elements")
	}

	// case2: set has one element
	s = New()
	s.Add(1)
	s.Remove(1)
	if s.Size() != 0 || s.Contains(1) {
		t.Error("case2 error: set has one element")
	}

	// case3: set has no element
	s = New()
	s.Remove(1)
	s.Remove(2)
	s.Remove(3)
	if s.Size() != 0 || s.Contains(1) || s.Contains(2) || s.Contains(3) {
		t.Error("case3 error: set has no element")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: set has some elements
	s := New()
	s.Add(1, 2, 3, 4)
	if s.Empty() {
		t.Error("case1 error: set has some elements")
	}

	// case2: set has one element
	s = New()
	s.Add(1)
	if s.Empty() {
		t.Error("case2 error: set has one element")
	}

	// case3: set has no element
	s = New()
	if !s.Empty() {
		t.Error("case3 error: set has no element")
	}
}

func TestSize(t *testing.T) {
	// case1: set has some elements
	s := New()
	s.Add(1, 2, 3, 4)
	if s.Size() != 4 {
		t.Error("case1 error: set has some elements")
	}

	// case2: set has one element
	s = New()
	s.Add(1)
	if s.Size() != 1 {
		t.Error("case2 error: set has one element")
	}

	// case3: set has no element
	s = New()
	if s.Size() != 0 {
		t.Error("case3 error: set has no element")
	}
}

func TestClear(t *testing.T) {
	// case1: set has some elements
	s := New()
	s.Add(1, 2, 3, 4)
	s.Clear()
	if s.Size() != 0 {
		t.Error("case1 error: set has some elements")
	}

	// case2: set has one element
	s = New()
	s.Add(1)
	s.Clear()
	if s.Size() != 0 {
		t.Error("case2 error: set has one element")
	}

	// case3: set has no element
	s = New()
	s.Clear()
	if s.Size() != 0 {
		t.Error("case3 error: set has no element")
	}
}

func TestValuesAndContains(t *testing.T) {
	// case1: set has some elements
	s := New()
	s.Add(1, 2, 3, 4)
	values := s.Values()
	if !s.Contains(values...) {
		t.Error("case1 error: set has some elements")
	}

	// case2: set has one element
	s = New()
	s.Add(1)
	values = s.Values()
	if !s.Contains(values...) {
		t.Error("case2 error: set has one element")
	}

	// case3: set has no element
	s = New()
	values = s.Values()
	if !s.Contains(values...) {
		t.Error("case3 error: set has no element")
	}
}
