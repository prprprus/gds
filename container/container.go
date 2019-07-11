package container

// Container Interface
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}
