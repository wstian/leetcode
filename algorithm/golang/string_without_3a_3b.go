// Given two integers A and B, return any string S such that:
//
// S has length A + B and contains exactly A 'a' letters, and exactly B 'b' letters;
// The substring 'aaa' does not occur in S;
// The substring 'bbb' does not occur in S.
//
//
// Example 1:
//
// Input: A = 1, B = 2
// Output: "abb"
// Explanation: "abb", "bab" and "bba" are all correct answers.
//
// Example 2:
//
// Input: A = 4, B = 1
// Output: "aabaa"

package main

import "strings"

func strWithout3a3b(A int, B int) string {
    if A == 0 {
        return strings.Repeat("b", B)
    }
    if B == 0 {
        return strings.Repeat("a", A)
    }
    if A > B {
        return "aab" + strWithout3a3b(A - 2, B -1)
    }
    if B > A {
        return strWithout3a3b(A - 1, B - 2)  + "abb"
    }
    return strings.Repeat("ab", A)
}