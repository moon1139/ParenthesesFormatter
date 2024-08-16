package main

import (
	"testing"
)

func TestFormatParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"(A*(B+C))", "A*(B+C)"},
		{"2+(3/-5)", "2+3/-5"},
		{"x+(y+z)+(t+(v+w))", "x+y+z+t+v+w"},
		{"((A+B))", "A+B"},
		{"((A*(B+C)))", "A*(B+C)"},
		{"(A+B)*(C+D)", "(A+B)*(C+D)"},
		{"((A+B)*(C+D))", "(A+B)*(C+D)"},
		{"(A+(B*(C+D)))", "A+B*(C+D)"},
		{"((A+B)+C)", "A+B+C"},
		{"((A+B)+(C+D))", "A+B+C+D"},
	}

	for _, test := range tests {
		result := formatParentheses(test.input)
		if result != test.expected {
			t.Errorf("formatParentheses(%q);\n got %q, want %q", test.input, result, test.expected)
		}
	}
}