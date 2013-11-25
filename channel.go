package karate

type channelChopper struct {}

func ChannelChopper() Chopper {
	return &channelChopper{}
}

func (*channelChopper) Chop(needle int, haystack []int) int {
	index := make(chan int)

	go channelChop(needle, haystack, index, 0, len(haystack)-1)

	return <-index
}

func channelChop(needle int, haystack []int, index chan int, low, high int) {
	if low > high {
		index <- -1

		return
	}

	middle := low + ((high - low) / 2)

	switch item := haystack[middle]; {
	case item == needle:
		index <- middle

		return
	case item > needle:
		high = middle - 1
	case item < needle:
		low = middle + 1
	}

	go channelChop(needle, haystack, index, low, high)
}
