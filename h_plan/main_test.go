package main

import (
	"testing"
)

func TestXxx(t *testing.T) {
	tree := NewTree(func(a, b Item) int { return a.(int) - b.(int) })
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(10)

	var iter Iterator
	iter = tree.Min()
	min := iter.Item()
	iter = tree.Max()
	max := iter.Item()

	tree.DeleteWithKey(10)

	iter = tree.Max()
	max = 0
	max = iter.Item()

	if 0 == min || 0 == max {
		t.Error()
	}
}
