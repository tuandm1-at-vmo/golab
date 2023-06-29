package lab

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	text := "This is a sample text. This text is just an example."
	expect := map[string]int{
		"just":    1,
		"This":    2,
		"is":      2,
		"a":       1,
		"sample":  1,
		"text":    2,
		"an":      1,
		"example": 1,
	}
	testCountWords(t, text, expect)
}

func testCountWords(t *testing.T, text string, expect map[string]int) {
	if actual := CountWords(text); !reflect.DeepEqual(actual, expect) {
		t.Errorf("failed: expect=%#v actual=%#v", expect, actual)
	}
}
