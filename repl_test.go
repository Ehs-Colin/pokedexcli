package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This is a test",
			expected: []string{"this", "is", "a", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}

func TestAttemptCatch(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{
		{
			input:    0,
			expected: true,
		},
		{
			input:    301,
			expected: false,
		},
	}
	for _, c := range cases {
		actual := attemptCatch(c.input)
		if actual != c.expected {
			t.Errorf("attemptCatch(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}

	trueAchieved := false
	falseAchieved := true
	for i := range 1000 {
		result := attemptCatch(150)
		if result {
			trueAchieved = true
		} else {
			falseAchieved = true
		}
		if trueAchieved && falseAchieved {
			fmt.Printf("Both results achieved in %d tries", i)
			break
		}
	}
}
