package skiplist

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	// DEBUG_DEFAULTLEVEL = 4
	// DEBUG_P = 1
	DEFAULTLEVEL = 32
	P            = 0.65
)

var (
	ErrNotFound = errors.New("can not found value")
	ErrEmpty    = errors.New("skiplist is empty")
)

type node struct {
	// key   interface{}
	// value interface{}
	key   int
	next  []*node
	level int
}

// SkipList
type SkipList struct {
	head     *node
	size     int
	maxLevel int
	// comparator util.Comparator
}

type record struct {
	currNode *node // the node that bigger than value
	rindex   int   // the index that node in next array position
}

// New
func New() *SkipList {
	skiplist := &SkipList{
		head:     new(node),
		size:     0,
		maxLevel: DEFAULTLEVEL,
	}
	skiplist.head.next = make([]*node, DEFAULTLEVEL)
	return skiplist
}

// Skiplist Interface

// Returns a new random level.
func randomLevel() (n int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	return rand.Intn(DEFAULTLEVEL-min) + min
}

// Set
func (skiplist *SkipList) Set(key int) {
	// find location of insertion
	recordArray := skiplist.find(key)

	// new Node
	p := randomLevel()
	// p := DEBUG_P
	newNode := &node{
		key:   key,
		next:  make([]*node, p),
		level: p,
	}

	// insert
	for i := 0; i < p; i++ {
		currNode := recordArray[i].currNode
		rindex := recordArray[i].rindex
		newNode.next[i] = currNode.next[rindex]
		currNode.next[rindex] = newNode
	}

	skiplist.size++
}

// find
func (skiplist *SkipList) find(key int) []record {
	recordArray := make([]record, DEFAULTLEVEL)
	currNode := skiplist.head

x:
	for i := len(currNode.next) - 1; i >= 0; {
		// CASE1-1: move down
		if currNode.next[i] == nil {
			recordArray = addRecordArray(recordArray, currNode, i)
			i--
			continue
		}

		// CASE2: move right
		for key > currNode.next[i].key {
			currNode = currNode.next[i]
			goto x
		}

		// CASE1-2: move down
		if key <= currNode.next[i].key {
			recordArray = addRecordArray(recordArray, currNode, i)
			i--
		}
	}

	return recordArray
}

func addRecordArray(recordArray []record, currNode *node, rindex int) []record {
	record := record{
		currNode: currNode,
		rindex:   rindex,
	}
	recordArray[rindex] = record
	return recordArray
}

// Show
func (skiplist *SkipList) Show() {
	flag := skiplist.head.next[0]
	for flag != nil {
		fmt.Println(flag.key)
		flag = flag.next[0]
	}
}

// Exists
func (skiplist *SkipList) Exists(key int) bool {
	recordArray := skiplist.find(key)

	// can not find
	if recordArray[0].currNode.next[0] == nil || recordArray[0].currNode.next[0].key != key {
		return false
	}

	return true
}

// Remove
func (skiplist *SkipList) Remove(key int) error {
	if skiplist.size == 0 {
		return ErrEmpty
	}

	recordArray := skiplist.find(key)

	// can not find
	if recordArray[0].currNode.next[0] == nil || recordArray[0].currNode.next[0].key != key {
		return ErrNotFound
	}

	// remove(adjustment pointer)
	level := recordArray[0].currNode.next[0].level
	for i := 0; i < level; i++ {
		recordArray[i].currNode.next[i] = recordArray[i].currNode.next[i].next[i]
	}
	skiplist.size--
	return nil
}

// Container Interface

// Empty returns true if the skiplist is empty, otherwise returns false.
func (skiplist *SkipList) Empty() bool {
	return skiplist.Size() == 0
}

// Size returns the size of the skiplist.
func (skiplist *SkipList) Size() int {
	return skiplist.size
}

// Clear the skiplist.
func (skiplist *SkipList) Clear() {
	skiplist.head = nil
	skiplist.size = 0
	skiplist.maxLevel = 0
}

// Values returns the values of skiplist.
func (skiplist *SkipList) Values() []interface{} {
	values := make([]interface{}, 0)
	flag := skiplist.head.next[0]
	for flag != nil {
		values = append(values, flag.key)
		flag = flag.next[0]
	}
	return values
}
