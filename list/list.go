package list

// List interface
type List interface {
	Append(values ...interface{})
	PreAppend(values ...interface{})
	indexInRange(index int) bool
	Get(index int) (interface{}, error)
	Remove(index int) error
	Contains(values ...interface{}) bool
	Swap(i, j int) error
	Insert(index int, values ...interface{}) error
	Set(index int, value interface{}) error
	IndexOf(value interface{}) (int, error)
	Reverse()
}
