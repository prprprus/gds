// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Package skiplist implements the skip list.
// Structure is not concurrent safe.
// Reference: https://en.wikipedia.org/wiki/Skip_list
package skiplist

import (
	"errors"
	"math/rand"
	"time"

	"github.com/prprprus/ds/util"
)

const (
	// DefaultMaxLevel is the default maximum height of the skip list
	DefaultMaxLevel = 32
)

var (
	// ErrNotFound is returned when the value can not found
	ErrNotFound = errors.New("can not found value")
	// ErrEmpty is returned when the skip list is empty
	ErrEmpty = errors.New("skiplist is empty")
)

// Node of the skip list.
type node struct {
	key   interface{} // value sorts by key
	value interface{} // value fields can store any type
	next  []*node     // next array of the skip list, pointing to multiple node struct behind
	level int         // height of node
}

// SkipList represents a skip list structure.
type SkipList struct {
	head       *node           // head is a `[]*node`, length equal to DefaultMaxLevel
	size       int             // number of nodes of the skip list
	maxLevel   int             // height of the skip list
	comparator util.Comparator // comparator is used to compare the size of the key
}

// Node of the path array.
type pathNode struct {
	pNode  *node // the node that bigger than value
	pIndex int   // the index that node in next array position
}

// New the skip list.
func New(comparator func(a, b interface{}) int) *SkipList {
	skiplist := &SkipList{
		head:       new(node),
		size:       0,
		maxLevel:   DefaultMaxLevel,
		comparator: comparator,
	}
	skiplist.head.next = make([]*node, DefaultMaxLevel)
	return skiplist
}

// Skiplist Interface

// Returns a new random level.
func randomLevel() (n int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	return rand.Intn(DefaultMaxLevel-min) + min
}

// Set key-value in the skip list.
func (skiplist *SkipList) Set(key, value interface{}) {
	path := skiplist.find(key)

	level := randomLevel()
	newNode := &node{
		key:   key,
		value: value,
		next:  make([]*node, level),
		level: level,
	}

	// insert according to level
	for i := 0; i < level; i++ {
		pNode := path[i].pNode
		pIndex := path[i].pIndex
		newNode.next[i] = pNode.next[pIndex]
		pNode.next[pIndex] = newNode
	}
	skiplist.size++
}

// Find the traversal path of the skip list.
//
// Divided into three cases:
// case1:
// 	If the next node of the current node is nil,
// 	then think that the key is smaller than the next node,
// 	record the current node to the path array and the current node moves down the next array.
// case2:
// 	If the key is larger than the next node of the current node,
// 	record the current node to the path array and the current node moves to the right along the next node.
// case3:
// 	If the key is less than or equal to the next node of the current node,
//	record the current node to the path array and the current node moves down the next array.
func (skiplist *SkipList) find(key interface{}) []pathNode {
	path := make([]pathNode, DefaultMaxLevel)
	currNode := skiplist.head

x:
	for i := len(currNode.next) - 1; i >= 0; {
		// case1: move down
		if currNode.next[i] == nil {
			path = addPath(path, currNode, i)
			i--
			continue
		}

		// case2: move right
		for skiplist.comparator(key, currNode.next[i].key) == 1 {
			currNode = currNode.next[i]
			// Note: `goto` should be used instead of continue or break.
			// The difference between them can be referred to
			// https://medium.com/golangspec/labels-in-go-4ffd81932339.
			goto x
		}

		// case3: move down
		if skiplist.comparator(key, currNode.next[i].key) == 0 || skiplist.comparator(key, currNode.next[i].key) == -1 {
			path = addPath(path, currNode, i)
			i--
		}
	}

	return path
}

// Add the current node into path array.
func addPath(path []pathNode, currNode *node, index int) []pathNode {
	record := pathNode{
		pNode:  currNode,
		pIndex: index,
	}
	path[index] = record
	return path
}

// Exists returns true if the key is exists, otherwise returns false.
func (skiplist *SkipList) Exists(key interface{}) bool {
	path := skiplist.find(key)

	// can't found
	if path[0].pNode.next[0] == nil || path[0].pNode.next[0].key != key {
		return false
	}

	return true
}

// Remove the value by key.
func (skiplist *SkipList) Remove(key interface{}) error {
	if skiplist.size == 0 {
		return ErrEmpty
	}

	path := skiplist.find(key)

	// can't found
	if path[0].pNode.next[0] == nil || path[0].pNode.next[0].key != key {
		return ErrNotFound
	}

	// remove according to level
	level := path[0].pNode.next[0].level
	for i := 0; i < level; i++ {
		path[i].pNode.next[i] = path[i].pNode.next[i].next[i]
	}
	skiplist.size--
	return nil
}

// Check if the index is within the length of the skiplist.
func (skiplist *SkipList) indexInRange(index int) bool {
	if index >= 0 && index < skiplist.size {
		return true
	}
	return false
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
