package zset

import (
	"math/rand"
)

// SkiplistMaxLevel 最大leavel
const SkiplistMaxLevel = 32

// SkipListBranch 分支
const SkipListBranch = 4

type skiplistLevel struct {
	forward *Element
	span    int
}

// Element 储存单元
type Element struct {
	Value    Interface
	backward *Element
	level    []*skiplistLevel
}

// Next returns the next skiplist element or nil.
func (e *Element) Next() *Element {
	return e.level[0].forward
}

// Prev returns the previous skiplist element of nil.
func (e *Element) Prev() *Element {
	return e.backward
}

// newElement returns an initialized element.
func newElement(level int, v Interface) *Element {
	slLevels := make([]*skiplistLevel, level)
	for i := 0; i < level; i++ {
		slLevels[i] = new(skiplistLevel)
	}

	return &Element{
		Value:    v,
		backward: nil,
		level:    slLevels,
	}
}

// randomLevel returns a random level.
func randomLevel() int {
	level := 1
	for (rand.Int31()&0xFFFF)%SkipListBranch == 0 {
		level++
	}

	if level < SkiplistMaxLevel {
		return level
	}
	return SkiplistMaxLevel
}
