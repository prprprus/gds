// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package stack

import "github.com/prprprus/ds/container"

// Stack Interface
type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
