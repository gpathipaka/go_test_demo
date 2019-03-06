package main

import "testing"

type testpair struct {
	x, y, result int
}

func TestSum(t *testing.T) {
	r := Sum(5, 6)

	if r != 11 {
		t.Error("Expected 12 got ", r)
	}

}

func TestSub(t *testing.T) {
	table := []testpair{
		{1, 4, -3},
		{8, 4, 4},
		{27, 15, 12},
		{10, 4, 6},
	}

	for _, pair := range table {
		result := Sub(pair.x, pair.y)
		if result != pair.result {
			t.Error("Expected ", pair.result, " And got ", result)
		}
	}
}
