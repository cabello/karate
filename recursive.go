package karate

import "math"

func RecursiveChop(needle int, haystack []int) int {
	index, found := recursiveChop(needle, haystack, 0)
	if !found {
		return -1
	}

	return index
}

func recursiveChop(needle int, haystack []int, extraOffset int) (index int, found bool) {
	if len(haystack) == 0 {
		return -1, false
	}

	offset := int(math.Floor(float64(len(haystack) / 2)))

	switch item := haystack[offset]; {
	case item == needle:
		return extraOffset + offset, true
	case item > needle:
		return recursiveChop(needle, haystack[:offset], extraOffset)
	case item < needle:
		offset++
		return recursiveChop(needle, haystack[offset:], offset)
	}

	return
}