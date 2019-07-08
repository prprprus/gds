package arrayqueue

import (
	"errors"

	"github.com/prprprus/ds/list/arraylist"
)

var (
	// ErrPop is returned when the queue is empty
	ErrPop = errors.New("queue is empty")
)

type Queue struct {
	list *arraylist.List
}

func New() *Queue {
	return &Queue{
		list: &arraylist.List{},
	}
}

func (queue *Queue) Put(value interface{}) {
	queue.list.Append(value)
}

func (queue *Queue) Get() (interface{}, error) {
	if queue.Size() == 0 {
		return nil, ErrPop
	}

	value, _ := queue.list.Get(0)
	queue.list.Remove(0)
	return value, nil
}

// Check if the index is within the length of the list.
func (queue *Queue) indexInRange(index int) bool {
	if index >= 0 && index < queue.list.Size() {
		return true
	}
	return false
}

// Container Interface

// Empty returns true if the queue is empty, otherwise returns false.
func (queue *Queue) Empty() bool {
	return queue.list.Size() == 0
}

// Size returns the size of the queue.
func (queue *Queue) Size() int {
	return queue.list.Size()
}

// Clear the queue.
func (queue *Queue) Clear() {
	queue.list.Clear()
}

// Values returns the values of queue.
func (queue *Queue) Values() []interface{} {
	return queue.list.Values()
}
