package karate

import (
	"reflect"
	"testing"
)

func TestAllChoppers(t *testing.T) {
	choppers := []Chopper{
		ChannelChopper(),
		IterativeChopper(),
		RecursiveChopper(),
		TraditionalChopper(),
		TreeChopper(),
	}

	for _, chopper := range choppers {
		testElementNotFound(t, chopper)
		testElementFoundExtremeLeft(t, chopper)
		testElementFoundExtremeRight(t, chopper)
		testElementFoundAnywhere(t, chopper)
		testElementFirstRightSecondLeft(t, chopper)
	}
}

func testElementNotFound(t *testing.T, chopper Chopper) {
	assertChop(t, -1, chopper, 3, []int{})
	assertChop(t, -1, chopper, 3, []int{1})
	assertChop(t, -1, chopper, 0, []int{1, 3, 5})
	assertChop(t, -1, chopper, 2, []int{1, 3, 5})
	assertChop(t, -1, chopper, 4, []int{1, 3, 5})
	assertChop(t, -1, chopper, 6, []int{1, 3, 5})
	assertChop(t, -1, chopper, 0, []int{1, 3, 5, 7})
	assertChop(t, -1, chopper, 2, []int{1, 3, 5, 7})
	assertChop(t, -1, chopper, 4, []int{1, 3, 5, 7})
	assertChop(t, -1, chopper, 6, []int{1, 3, 5, 7})
	assertChop(t, -1, chopper, 8, []int{1, 3, 5, 7})
}

func testElementFoundExtremeLeft(t *testing.T, chopper Chopper) {
	assertChop(t, 0, chopper, 1, []int{1})
	assertChop(t, 0, chopper, 1, []int{1, 3, 5})
	assertChop(t, 0, chopper, 1, []int{1, 3, 5, 7})
}

func testElementFoundExtremeRight(t *testing.T, chopper Chopper) {
	assertChop(t, 2, chopper, 5, []int{1, 3, 5})
	assertChop(t, 3, chopper, 7, []int{1, 3, 5, 7})
}

func testElementFoundAnywhere(t *testing.T, chopper Chopper) {
	assertChop(t, 1, chopper, 3, []int{1, 3, 5})
	assertChop(t, 1, chopper, 3, []int{1, 3, 5, 7})
	assertChop(t, 2, chopper, 5, []int{1, 3, 5, 7})
}

func testElementFirstRightSecondLeft(t *testing.T, chopper Chopper) {
	assertChop(t, 7, chopper, 12, []int{1, 3, 4, 5, 7, 9, 10, 12, 14, 15, 16})
}

func assertChop(t *testing.T, expected int, chopper Chopper, needle int, haystack []int) {
	if result := chopper.Chop(needle, haystack); result != expected {
		t.Errorf("%s(%v, %v) = %v, want %v", getObjectName(chopper), needle, haystack, result, expected)
	}
}

func getObjectName(i interface{}) string {
	return reflect.TypeOf(i).String()
}
