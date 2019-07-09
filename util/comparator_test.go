// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package util

import "testing"

// int, int8, int16, int32, int64

func TestIntComparator(t *testing.T) {
	// case1: a < b
	a := 1
	b := 2
	if IntComparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = 2
	b = 1
	if IntComparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = 1
	b = 1
	if IntComparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestInt8Comparator(t *testing.T) {
	// case1: a < b
	a := int8(123)
	b := int8(125)
	if Int8Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = int8(125)
	b = int8(123)
	if Int8Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = int8(123)
	b = int8(123)
	if Int8Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestInt16Comparator(t *testing.T) {
	// case1: a < b
	a := int16(32761)
	b := int16(32763)
	if Int16Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = int16(32763)
	b = int16(32761)
	if Int16Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = int16(32763)
	b = int16(32763)
	if Int16Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestInt32Comparator(t *testing.T) {
	// case1: a < b
	a := int32(2147483641)
	b := int32(2147483644)
	if Int32Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = int32(2147483644)
	b = int32(2147483641)
	if Int32Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = int32(2147483644)
	b = int32(2147483644)
	if Int32Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestInt64Comparator(t *testing.T) {
	// case1: a < b
	a := int64(9223372036854775802)
	b := int64(9223372036854775805)
	if Int64Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = int64(9223372036854775805)
	b = int64(9223372036854775802)
	if Int64Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = int64(9223372036854775802)
	b = int64(9223372036854775802)
	if Int64Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

// uint, uint8, uint16, uint32, uint64

func TestUIntComparator(t *testing.T) {
	// case1: a < b
	a := uint(12)
	b := uint(22)
	if UIntComparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = uint(22)
	b = uint(12)
	if UIntComparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = uint(12)
	b = uint(12)
	if UIntComparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestUInt8Comparator(t *testing.T) {
	// case1: a < b
	a := uint8(252)
	b := uint8(253)
	if UInt8Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = uint8(253)
	b = uint8(252)
	if UInt8Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = uint8(253)
	b = uint8(253)
	if UInt8Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}
