package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abba", true},
		{"abcba", true},
		{"abcdcba", true},
		{"abcde", false},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
