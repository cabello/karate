package karate

import "math"

type iterativeChopper struct {}

func IterativeChopper() Chopper {
	return &iterativeChopper{}
}

func (*iterativeChopper) Chop(needle int, haystack []int) int {
	extraOffset := 0

	for {
		if len(haystack) == 0 {
			return -1
		}

		offset := int(math.Floor(float64(len(haystack) / 2)))

		switch item := haystack[offset]; {
		case item == needle:
			return extraOffset + offset
		case item > needle:
			haystack = haystack[:offset]
		case item < needle:
			offset++
			extraOffset += offset
			haystack = haystack[offset:]
		}
	}
}
