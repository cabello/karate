package karate

import "math"

type Tree struct {
	index int
	value int
	left  *Tree
	right *Tree
}

func buildTree(haystack []int, extraOffset int) *Tree {
	if len(haystack) == 0 {
		return nil
	}

	offset := int(math.Floor(float64(len(haystack) / 2)))

	tree := &Tree{
		index: extraOffset + offset,
		value: haystack[offset],
	}

	tree.left = buildTree(haystack[:offset], extraOffset)
	offset++
	tree.right = buildTree(haystack[offset:], extraOffset+offset)

	return tree
}

func TreeChop(needle int, haystack []int) int {
	tree := buildTree(haystack, 0)

	for {
		if tree == nil {
			return -1
		} else if tree.value == needle {
			return tree.index
		} else if tree.value < needle {
			tree = tree.right
		} else if tree.value > needle {
			tree = tree.left
		}
	}
}
