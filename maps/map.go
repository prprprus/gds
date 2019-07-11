// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package maps

import "github.com/prprprus/ds/container"

// Map Interface
type Map interface {
	Put(key, value interface{})
	Get(key interface{}) (interface{}, error)
	Remove(key interface{})

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
