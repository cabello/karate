package karate

type TraditionalChopper struct {
	haystack []int
}

func (t *TraditionalChopper) Init(haystack []int) {
	t.haystack = haystack
}

func (t *TraditionalChopper) Chop(needle int) int {
	low := 0
	high := len(t.haystack) - 1

	for high-low >= 0 {
		middle := low + ((high - low) / 2)

		switch item := t.haystack[middle]; {
		case item == needle:
			return middle
		case item > needle:
			high = middle - 1
		case item < needle:
			low = middle + 1
		}
	}

	return -1
}
