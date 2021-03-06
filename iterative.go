package karate

import "math"

type IterativeChopper struct {
	haystack []int
}

func (i *IterativeChopper) Init(haystack []int) {
	i.haystack = haystack
}

func (i *IterativeChopper) Chop(needle int) int {
	extraOffset := 0

	for {
		if len(i.haystack) == 0 {
			return -1
		}

		offset := int(math.Floor(float64(len(i.haystack) / 2)))

		switch item := i.haystack[offset]; {
		case item == needle:
			return extraOffset + offset
		case item > needle:
			i.haystack = i.haystack[:offset]
		case item < needle:
			offset++
			extraOffset += offset
			i.haystack = i.haystack[offset:]
		}
	}
}
