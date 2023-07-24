package mod01

import (
	"reflect"
	"testing"
)

func TestGetAllPrimesUpTo(t *testing.T) {
	var n uint = 100
	var expect = []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	testGetAllPrimesUpTo(t, n, expect)
}

func testGetAllPrimesUpTo(t *testing.T, n uint, expect []uint) {
	actual, err := GetAllPrimesUpTo(n, "./first-1000000-primes.json")
	if err != nil {
		t.Errorf("failed: %s", err.Error())
	}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("failed: expect=%#v actual=%#v", expect, actual)
	}
}
