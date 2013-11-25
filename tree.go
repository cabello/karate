package karate

import "math"

type TreeChopper struct {
	tree *tree
}

type tree struct {
	index int
	value int
	left  *tree
	right *tree
}

func (t *TreeChopper) Init(haystack []int) {
	t.tree = buildTree(haystack, 0)
}

func buildTree(haystack []int, extraOffset int) *tree {
	if len(haystack) == 0 {
		return nil
	}

	offset := int(math.Floor(float64(len(haystack) / 2)))

	tree := &tree{
		index: extraOffset + offset,
		value: haystack[offset],
	}

	tree.left = buildTree(haystack[:offset], extraOffset)
	offset++
	tree.right = buildTree(haystack[offset:], extraOffset+offset)

	return tree
}

func (t *TreeChopper) Chop(needle int) int {
	for {
		if t.tree == nil {
			return -1
		} else if t.tree.value == needle {
			return t.tree.index
		} else if t.tree.value < needle {
			t.tree = t.tree.right
		} else if t.tree.value > needle {
			t.tree = t.tree.left
		}
	}
}
