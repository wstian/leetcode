package main

import "testing"

func TestStringWithout3A3B(t *testing.T)  {
    cases := []struct{
        A int;
        B int;
        expected string;
    } {
        {1, 2, "abb"},
        {4, 1, "aabaa"},
    }
    for _, c := range cases {
        result := strWithout3a3b(c.A, c.B)
        if result != c.expected {
            t.Errorf("case: '%v', got: '%v'\n", c, result)
        }
    }
}
