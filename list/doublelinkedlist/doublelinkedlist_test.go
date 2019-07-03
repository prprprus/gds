package doublelinkedlist

import "testing"

func TestNew(t *testing.T) {
	// case1: new empty list
	list := New()
	if list.first != nil || list.last != nil || list.size != 0 {
		t.Error("case1 error: new empty list")
	}
}
