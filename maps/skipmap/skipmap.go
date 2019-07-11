package skipmap

import (
	"errors"

	"github.com/prprprus/ds/skiplist"
)

var (
	// ErrNotFound is returned when the value can not found
	ErrNotFound = errors.New("can not found value")
	// ErrEmpty is returned when the map is empty
	ErrEmpty = errors.New("the map is empty")
)

type Map struct {
	m *skiplist.SkipList
}
