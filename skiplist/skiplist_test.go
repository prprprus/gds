package skiplist

import (
	"testing"
)

func verifyElement(result []interface{}, skiplist *SkipList) bool {
	flag := skiplist.head.next[0]
	i := 0
	for flag != nil {
		if flag.value != result[i] {
			return false
		}
		flag = flag.next[0]
		i++
	}
	return true
}

func verifyValues(result, values []interface{}) bool {
	for i, v := range result {
		if values[i] != v {
			return false
		}
	}
	return true
}

func TestNew(t *testing.T) {
	skiplist := New()
	if skiplist.head == nil || skiplist.size != 0 || skiplist.maxLevel != DEFAULTLEVEL || len(skiplist.head.next) != DEFAULTLEVEL {
		t.Error("TestNew error")
	}
}

func TestRandomLevel(t *testing.T) {
	r := randomLevel()
	if r < 1 || randomLevel() > DEFAULTLEVEL {
		t.Error("TestRandomLevel error")
	}
}

// Skiplist Interface

// Note: cover `find` and `addRecordArray` method.
func TestSet(t *testing.T) {
	// case1: skiplist has some elements
	skiplist := New()
	skiplist.Set(777)
	skiplist.Set(999)
	skiplist.Set(99)
	skiplist.Set(1)
	skiplist.Set(12)
	skiplist.Set(3)
	skiplist.Set(42)
	skiplist.Set(9)
	skiplist.Set(15)
	skiplist.Set(6)
	skiplist.Set(79)
	skiplist.Set(18)
	skiplist.Set(63)
	skiplist.Set(81)
	skiplist.Set(-1)
	result := []interface{}{-1, 1, 3, 6, 9, 12, 15, 18, 42, 63, 79, 81, 99, 777, 999}
	if skiplist.size != 15 || !verifyElement(result, skiplist) {
		t.Error("case1 error: skiplist has some elements")
	}

	// case2: skiplist has one element
	skiplist = New()
	skiplist.Set(1)
	result = []interface{}{1}
	if skiplist.size != 1 || !verifyElement(result, skiplist) {
		t.Error("case2 error: skiplist has one element")
	}

	// case3: skiplist has no element
	skiplist = New()
	result = []interface{}{}
	if skiplist.size != 0 || !verifyElement(result, skiplist) {
		t.Error("case3 error: skiplist has no element")
	}
}

// Note: cover `find` and `addRecordArray` method.
func TestExists(t *testing.T) {
	// case1: skiplist has some elements
	skiplist := New()
	skiplist.Set(777)
	skiplist.Set(99)
	skiplist.Set(-1)
	if skiplist.size != 3 || !skiplist.Exists(777) || !skiplist.Exists(99) || !skiplist.Exists(-1) {
		t.Error("case1 error: skiplist has some elements")
	}

	// case2: skiplist has one element
	skiplist = New()
	skiplist.Set(1)
	if skiplist.size != 1 || !skiplist.Exists(1) {
		t.Error("case2 error: skiplist has one element")
	}

	// case3: skiplist has no element
	skiplist = New()
	if skiplist.size != 0 || skiplist.Exists(777) {
		t.Error("case3 error: skiplist has no element")
	}
}

// Note: cover `find` and `addRecordArray` method.
func TestRemove(t *testing.T) {
	// case1: skiplist has some elements
	skiplist := New()
	skiplist.Set(777)
	skiplist.Set(999)
	skiplist.Set(99)
	skiplist.Set(1)
	skiplist.Set(12)
	skiplist.Set(3)
	skiplist.Set(42)
	skiplist.Set(9)
	skiplist.Set(15)
	skiplist.Set(6)
	skiplist.Set(79)
	skiplist.Set(18)
	skiplist.Set(63)
	skiplist.Set(81)
	skiplist.Set(-1)
	skiplist.Remove(-1)
	skiplist.Remove(999)
	skiplist.Remove(79)
	skiplist.Remove(42)
	result := []interface{}{1, 3, 6, 9, 12, 15, 18, 63, 81, 99, 777}
	if skiplist.size != 11 || !verifyElement(result, skiplist) {
		t.Error("case1 error: skiplist has some elements")
	}

	// case2: skiplist has one element
	skiplist = New()
	skiplist.Set(1)
	skiplist.Remove(1)
	result = []interface{}{}
	if skiplist.size != 0 || !verifyElement(result, skiplist) {
		t.Error("case2 error: skiplist has one element")
	}

	// case3: skiplist has no element (including remove the values that do not exist)
	skiplist = New()
	skiplist.Remove(1)
	skiplist.Remove(2)
	skiplist.Remove(3)
	result = []interface{}{}
	if skiplist.size != 0 || !verifyElement(result, skiplist) {
		t.Error("case3 error: skiplist has no element (including remove the values that do not exist)")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: skiplist has some elements
	skiplist := New()
	skiplist.Set(777)
	skiplist.Set(999)
	skiplist.Set(99)
	if skiplist.Empty() {
		t.Error("case1 error: skiplist has some elements")
	}

	// case2: skiplist has one element
	skiplist = New()
	skiplist.Set(1)
	if skiplist.Empty() {
		t.Error("case2 error: skiplist has one element")
	}

	// case3: skiplist has no element
	skiplist = New()
	if !skiplist.Empty() {
		t.Error("case3 error: skiplist has no element")
	}
}

func TestSize(t *testing.T) {
	// case1: skiplist has some elements
	skiplist := New()
	skiplist.Set(777)
	skiplist.Set(999)
	skiplist.Set(99)
	if skiplist.size != 3 {
		t.Error("case1 error: skiplist has some elements")
	}

	// case2: skiplist has one element
	skiplist = New()
	skiplist.Set(1)
	if skiplist.Size() != 1 {
		t.Error("case2 error: skiplist has one element")
	}

	// case3: skiplist has no element
	skiplist = New()
	if skiplist.Size() != 0 {
		t.Error("case3 error: skiplist has no element")
	}
}

func TestClear(t *testing.T) {
	// case1: skiplist has some elements
	skiplist := New()
	skiplist.Set(777)
	skiplist.Set(999)
	skiplist.Set(99)
	skiplist.Clear()
	if skiplist.head != nil || skiplist.size != 0 || skiplist.maxLevel != 0 {
		t.Error("case1 error: skiplist has some elements")
	}

	// case2: skiplist has one element
	skiplist = New()
	skiplist.Set(1)
	skiplist.Clear()
	if skiplist.head != nil || skiplist.size != 0 || skiplist.maxLevel != 0 {
		t.Error("case2 error: skiplist has one element")
	}

	// case3: skiplist has no element
	skiplist = New()
	skiplist.Clear()
	if skiplist.head != nil || skiplist.size != 0 || skiplist.maxLevel != 0 {
		t.Error("case3 error: skiplist has no element")
	}
}

func TestValue(t *testing.T) {
	// case1: skiplist has some elements
	skiplist := New()
	skiplist.Set(777)
	skiplist.Set(999)
	skiplist.Set(99)
	skiplist.Set(1)
	skiplist.Set(12)
	skiplist.Set(3)
	skiplist.Set(42)
	skiplist.Set(9)
	skiplist.Set(15)
	skiplist.Set(6)
	skiplist.Set(79)
	skiplist.Set(18)
	skiplist.Set(63)
	skiplist.Set(81)
	skiplist.Set(-1)
	values := skiplist.Values()
	result := []interface{}{-1, 1, 3, 6, 9, 12, 15, 18, 42, 63, 79, 81, 99, 777, 999}
	if !verifyValues(result, values) {
		t.Error("case1 error: skiplist has some elements")
	}

	// case2: skiplist has one element
	skiplist = New()
	skiplist.Set(1)
	values = skiplist.Values()
	result = []interface{}{1}
	if !verifyValues(result, values) {
		t.Error("case2 error: skiplist has one element")
	}

	// case3: skiplist has no element
	skiplist = New()
	values = skiplist.Values()
	result = []interface{}{}
	if !verifyValues(result, values) {
		t.Error("case3 error: skiplist has no element")
	}
}
