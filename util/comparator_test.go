// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package util

import (
	"testing"
)

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

func TestUInt16Comparator(t *testing.T) {
	// case1: a < b
	a := uint16(65520)
	b := uint16(65533)
	if UInt16Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = uint16(65533)
	b = uint16(65520)
	if UInt16Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = uint16(65533)
	b = uint16(65533)
	if UInt16Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestUInt32Comparator(t *testing.T) {
	// case1: a < b
	a := uint32(4294967291)
	b := uint32(4294967294)
	if UInt32Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = uint32(4294967294)
	b = uint32(4294967291)
	if UInt32Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = uint32(4294967294)
	b = uint32(4294967294)
	if UInt32Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestUInt64Comparator(t *testing.T) {
	// case1: a < b
	a := uint64(18446744073709551611)
	b := uint64(18446744073709551613)
	if UInt64Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = uint64(18446744073709551613)
	b = uint64(18446744073709551611)
	if UInt64Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = uint64(18446744073709551613)
	b = uint64(18446744073709551613)
	if UInt64Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

// float32, float64

func TestFloat32Comparator(t *testing.T) {
	// case1: a < b
	a := float32(1.0000001190)
	b := float32(2.0000001192)
	if Float32Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = float32(2.0000001192)
	b = float32(1.0000001190)
	if Float32Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = float32(1.0000001192)
	b = float32(1.0000001192)
	if Float32Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestFloat64Comparator(t *testing.T) {
	// case1: a < b
	a := float64(1.0000000000000004)
	b := float64(2.0000000000000004)
	if Float64Comparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = float64(2.0000000000000004)
	b = float64(1.0000000000000004)
	if Float64Comparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = float64(2.0000000000000004)
	b = float64(2.0000000000000004)
	if Float64Comparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

// byte, rune, string

func TestByteComparator(t *testing.T) {
	// case1: a < b
	a := byte('a')
	b := byte('b')
	if ByteComparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = byte('b')
	b = byte('a')
	if ByteComparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = byte('b')
	b = byte('b')
	if ByteComparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestRuneComparator(t *testing.T) {
	// case1: a < b
	a := rune('a')
	b := rune('b')
	if RuneComparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = rune('b')
	b = rune('a')
	if RuneComparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = rune('b')
	b = rune('b')
	if RuneComparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}

func TestStringComparator(t *testing.T) {
	// case1: a < b
	a := string('a')
	b := string('b')
	if StringComparator(a, b) >= 0 {
		t.Error("case1 error: a < b")
	}

	// case2: a > b
	a = string('b')
	b = string('a')
	if StringComparator(a, b) <= 0 {
		t.Error("case2 error: a > b")
	}

	// case3: a == b
	a = string('b')
	b = string('b')
	if StringComparator(a, b) != 0 {
		t.Error("case3 error: a == b")
	}
}
