package karate

func TraditionalChop(needle int, haystack []int) int {
	low := 0
	high := len(haystack) - 1

	for high-low >= 0 {
		middle := low + ((high - low) / 2)

		switch item := haystack[middle]; {
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
