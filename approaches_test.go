package karate

import (
	"reflect"
	"runtime"
	"testing"
)

func TestAllApproaches(t *testing.T) {
	methods := []func(int, []int) int{
		ChannelChop,
		IterativeChop,
		RecursiveChop,
		TraditionalChop,
		TreeChop,
	}

	for _, method := range methods {
		testElementNotFound(t, method)
		testElementFoundExtremeLeft(t, method)
		testElementFoundExtremeRight(t, method)
		testElementFoundAnywhere(t, method)
		testElementFirstRightSecondLeft(t, method)
	}
}

func testElementNotFound(t *testing.T, chop func(int, []int) int) {
	assertChop(t, -1, chop, 3, []int{})
	assertChop(t, -1, chop, 3, []int{1})
	assertChop(t, -1, chop, 0, []int{1, 3, 5})
	assertChop(t, -1, chop, 2, []int{1, 3, 5})
	assertChop(t, -1, chop, 4, []int{1, 3, 5})
	assertChop(t, -1, chop, 6, []int{1, 3, 5})
	assertChop(t, -1, chop, 0, []int{1, 3, 5, 7})
	assertChop(t, -1, chop, 2, []int{1, 3, 5, 7})
	assertChop(t, -1, chop, 4, []int{1, 3, 5, 7})
	assertChop(t, -1, chop, 6, []int{1, 3, 5, 7})
	assertChop(t, -1, chop, 8, []int{1, 3, 5, 7})
}

func testElementFoundExtremeLeft(t *testing.T, chop func(int, []int) int) {
	assertChop(t, 0, chop, 1, []int{1})
	assertChop(t, 0, chop, 1, []int{1, 3, 5})
	assertChop(t, 0, chop, 1, []int{1, 3, 5, 7})
}

func testElementFoundExtremeRight(t *testing.T, chop func(int, []int) int) {
	assertChop(t, 2, chop, 5, []int{1, 3, 5})
	assertChop(t, 3, chop, 7, []int{1, 3, 5, 7})
}

func testElementFoundAnywhere(t *testing.T, chop func(int, []int) int) {
	assertChop(t, 1, chop, 3, []int{1, 3, 5})
	assertChop(t, 1, chop, 3, []int{1, 3, 5, 7})
	assertChop(t, 2, chop, 5, []int{1, 3, 5, 7})
}

func testElementFirstRightSecondLeft(t *testing.T, chop func(int, []int) int) {
	assertChop(t, 7, chop, 12, []int{1, 3, 4, 5, 7, 9, 10, 12, 14, 15, 16})
}

func assertChop(t *testing.T, expected int, chop func(int, []int) int, needle int, haystack []int) {
	if result := chop(needle, haystack); result != expected {
		t.Errorf("%s(%v, %v) = %v, want %v", getFunctionName(chop), needle, haystack, result, expected)
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
