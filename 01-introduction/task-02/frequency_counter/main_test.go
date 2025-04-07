package main

import (
	"reflect"
	"testing"
)

func TestStripPunctuation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test!", "test"},
		{"hello, world!", "hello world"},
		{"123!", "123"},
		{"", ""},
	}

	for _, test := range tests {
		result := stripPunctuation(test.input)
		if result != test.expected {
			t.Errorf("stripPunctuation(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"hello world", map[string]int{"hello": 1, "world": 1}},
		{"hello hello world", map[string]int{"hello": 2, "world": 1}},
		{"", map[string]int{}},
	}

	for _, test := range tests {
		result := countWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("countWords(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
