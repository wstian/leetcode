package main

import "testing"

func TestTwoCityScheduling(t *testing.T) {
	test_cases := []struct{
		costs [][]int;
		expected int;
	}{
		{[][]int{{10,20},{30,200},{400,50},{30,20}}, 110},
	}
	for _, c := range test_cases {
		result := twoCitySchedCost(c.costs)
		if result != c.expected {
			t.Errorf("case: '%v', got: '%v'\n", c, result)
		}
	}
}
