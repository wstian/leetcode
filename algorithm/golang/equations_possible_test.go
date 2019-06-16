package main

import "testing"

func TestEqualityEquations(t *testing.T) {
	test_cases := []struct{
		equations []string;
		expected bool;
	}{
		{[]string{"a==b","b!=c","c==a"}, false},
        {[]string{"a==b","b!=a"}, false},
        {[]string{"b==a","a==b"}, true},
		{[]string{"a==b","b!=c","c==a"}, false},
		{[]string{"c==c","b==d","x!=z"}, true},
	}
	for _, c := range test_cases {
		result := equationsPossible(c.equations)
		if result != c.expected {
			t.Errorf("case: '%v', got: '%v'\n", c, result)
		}
	}
}
