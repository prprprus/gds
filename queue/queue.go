// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package queue

import "github.com/prprprus/ds/container"

// Queue Interface
type Queue interface {
	Put(value interface{})
	Get() (interface{}, error)

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
