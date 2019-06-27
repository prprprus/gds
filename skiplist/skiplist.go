package skiplist

import (
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

// var (
// 	recordArray = make([]record, 32)
// )

type node struct {
	value int
	next  []*node
}

// SkipList
type SkipList struct {
	head     *node
	Size     int
	MaxLevel int
	top      *node
}

type record struct {
	currNode *node // 比 value 大的节点
	rindex   int   // 这个节点在 next 数组中的下标
}

// New
func New() *SkipList {
	skiplist := &SkipList{
		head:     new(node),
		Size:     0,
		MaxLevel: DEFAULTLEVEL,
	}
	skiplist.head.next = make([]*node, DEFAULTLEVEL)

	return skiplist
}

// Returns a new random level.
func randomLevel() (n int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	return rand.Intn(DEFAULTLEVEL-min) + min
}

// Set
func (skiplist *SkipList) Set(value int) {
	// find location of insertion
	recordArray := skiplist.find(value)

	// new Node
	p := randomLevel()
	// p := DEBUG_P
	newNode := &node{
		value: value,
		next:  make([]*node, p),
	}
	// fmt.Println("recordArray 长度：", len(recordArray))
	// fmt.Println("recordArray：", recordArray)
	// fmt.Println("recordArray 0：", recordArray[0].currNode.value)
	// fmt.Println("骰子：", p)

	// insert
	for i := 0; i < p; i++ {
		currNode := recordArray[i].currNode
		rindex := recordArray[i].rindex
		// fmt.Println("--->")
		// fmt.Println("len:", len(currNode.next))
		// fmt.Println("--->")
		newNode.next[i] = currNode.next[rindex]
		currNode.next[rindex] = newNode
	}

	skiplist.Size++
}

// find 找 value 的前一个节点（也就是比 value 小的节点中最大的一个）
func (skiplist *SkipList) find(value int) []record {
	recordArray := make([]record, DEFAULTLEVEL)
	currNode := skiplist.head

x:
	for i := len(currNode.next) - 1; i >= 0; {
		// move down
		if currNode.next[i] == nil {
			recordArray = addRecordArray(recordArray, currNode, i)
			i-- // 相当于指针下移
			continue
		}

		// move right
		for value > currNode.next[i].value {
			currNode = currNode.next[i]
			continue x
		}

		// move down
		if value <= currNode.next[i].value {
			recordArray = addRecordArray(recordArray, currNode, i)
			i--
		}
	}

	return recordArray
}

func addRecordArray(recordArray []record, currNode *node, rindex int) []record {
	// 记录中间变量
	record := record{
		currNode: currNode,
		rindex:   rindex,
	}
	recordArray[rindex] = record
	return recordArray
}

func (skiplist *SkipList) Show() {
	flag := skiplist.head.next[0]
	for flag != nil {
		fmt.Println(flag.value)
		flag = flag.next[0]
	}
}

// Get
func (skiplist *SkipList) Get(value int) bool {
	recordArray := skiplist.find(value)

	if recordArray[0].currNode.next[0] == nil {
		return false
	}

	if recordArray[0].currNode.next[0].value == value {
		return true

	}
	return false
}
