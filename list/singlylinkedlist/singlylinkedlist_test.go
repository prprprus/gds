package singlylinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New()
	if list.frist != nil || list.last != nil || list.size != 0 {
		t.Error("empty list initial error")
	}

	list = New(1, 2, 3, 4, 5)
	if list.size != 5 {
		t.Error("list initial error")
	}
}

func TestAppend(t *testing.T) {

}
