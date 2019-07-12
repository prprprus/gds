package set

import "github.com/prprprus/ds/container"

// Set Interface
type Set interface {
	Add(values ...interface{})
	Remove(values ...interface{}) error
	Contains(values ...interface{}) bool

	container.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
