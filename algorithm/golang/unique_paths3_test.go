package main

import "testing"

func TestUniquePaths3(t *testing.T) {
	test_cases := []struct{
		grid [][]int;
		expected int;
	}{
		{[][]int{{1,0,0,0},{0,0,0,0},{0,0,2,-1}}, 2},
	}
	for _, c := range test_cases {
		result := uniquePathsIII(c.grid)
		if result != c.expected {
			t.Errorf("case: '%v', got: '%v'\n", c, result)
		}
	}
}
