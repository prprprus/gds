// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package util

import (
	"strings"
	"time"
)

// Comparator is used for comparison between the same types.
//
// The types of built-in support are:
// 	int, int8, int16, int32, int64
// 	uint, uint8, uint16, uint32, uint64
// 	float32, float64
// 	byte, run, string
// 	time
//
// You can also use a custom comparator.
type Comparator func(a, b interface{}) int

// int, int8, int16, int32, int64

// IntComparator compares the value of type int.
func IntComparator(a, b interface{}) int {
	c1 := a.(int)
	c2 := b.(int)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// Int8Comparator compares the value of type int8.
func Int8Comparator(a, b interface{}) int {
	c1 := a.(int8)
	c2 := b.(int8)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// Int16Comparator compares the value of type int16.
func Int16Comparator(a, b interface{}) int {
	c1 := a.(int16)
	c2 := b.(int16)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// Int32Comparator compares the value of type int32.
func Int32Comparator(a, b interface{}) int {
	c1 := a.(int32)
	c2 := b.(int32)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// Int64Comparator compares the value of type int64.
func Int64Comparator(a, b interface{}) int {
	c1 := a.(int64)
	c2 := b.(int64)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// uint, uint8, uint16, uint32, uint64

// UIntComparator compares the value of type uint.
func UIntComparator(a, b interface{}) int {
	c1 := a.(uint)
	c2 := b.(uint)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// UInt8Comparator compares the value of type uint8.
func UInt8Comparator(a, b interface{}) int {
	c1 := a.(uint8)
	c2 := b.(uint8)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// UInt16Comparator compares the value of type uint16.
func UInt16Comparator(a, b interface{}) int {
	c1 := a.(uint16)
	c2 := b.(uint16)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// UInt32Comparator compares the value of type uint32.
func UInt32Comparator(a, b interface{}) int {
	c1 := a.(uint32)
	c2 := b.(uint32)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// UInt64Comparator compares the value of type uint64.
func UInt64Comparator(a, b interface{}) int {
	c1 := a.(uint64)
	c2 := b.(uint64)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// float32, float64

// Float32Comparator compares the value of type float32.
func Float32Comparator(a, b interface{}) int {
	c1 := a.(float32)
	c2 := b.(float32)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// Float64Comparator compares the value of type float64.
func Float64Comparator(a, b interface{}) int {
	c1 := a.(float64)
	c2 := b.(float64)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// byte, rune, string

// ByteComparator compares the value of type byte.
func ByteComparator(a, b interface{}) int {
	c1 := a.(byte)
	c2 := b.(byte)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// RuneComparator compares the value of type rune.
func RuneComparator(a, b interface{}) int {
	c1 := a.(rune)
	c2 := b.(rune)
	switch {
	case c1 < c2:
		return -1
	case c1 > c2:
		return 1
	default:
		return 0
	}
}

// StringComparator compares the value of type string.
func StringComparator(a, b interface{}) int {
	c1 := a.(string)
	c2 := b.(string)
	return strings.Compare(c1, c2)
}

// time

// TimeCompare compares the value of type time.
func TimeCompare(a, b interface{}) int {
	c1 := a.(time.Time)
	c2 := b.(time.Time)
	switch {
	case c1.Before(c2):
		return -1
	case c1.After(c2):
		return 1
	default:
		return 0
	}
}
