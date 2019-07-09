// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package util

import "testing"

func TestIntComparator(t *testing.T) {
	// case1: a == b
	a := 1
	b := 1
	if IntComparator(a, b) != 0 {
		t.Error("case1 error: a == b")
	}
}
