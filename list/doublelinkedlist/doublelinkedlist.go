package doublelinkedlist

// List
type List struct {
	frist *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	prev  *element
	next  *element
}

// New
func New(value ...interface{}) *List {
	list := &List{}

	if len(value) > 0 {
		list.Append(value...)
	}

	return list
}

// List Interface

// Append values (one or more than one) to list.
func (list *List) Append(value ...interface{}) {
	if len(value) == 0 {
		return
	}

	// if size is equal to 0, it is a new double linked list
	if list.size == 0 {
		for i, v := range value {
			newElement := &element{value: v}
			// first element
			if i == 0 {
				list.frist = newElement
				list.last = newElement
			} else {
				list.last.next = newElement
				newElement.prev = list.last
				list.last = newElement
			}
			list.size++
		}
	} else {
		for _, v := range value {
			newElement := &element{value: v}
			list.last.next = newElement
			newElement.prev = list.last
			list.last = newElement
			list.size++
		}
	}
}

// Check if the index is within the length of the list
func (list *List) indexInRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}
