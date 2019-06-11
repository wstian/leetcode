package main

import "testing"

func TestBrokenCalc(t *testing.T) {
    test_cases := []struct{
		X int;
		Y int;
		expected int;
	}{
		{2, 3, 2},
		{5, 8, 2},
		{3, 10, 3},
	}
	for _, c := range test_cases {
		result := brokenCalc(c.X, c.Y)
		if result != c.expected {
			t.Errorf("case: '%v', got: '%v'\n", c, result)
		}
	}
}
