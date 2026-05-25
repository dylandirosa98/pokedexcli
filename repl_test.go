package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "     hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hey whats going on dude",
			expected: []string{"hey", "whats", "going", "on", "dude"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of cleanInput(%v) = %d; want %d", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word[%v] = %s; want %s", i, word, expectedWord)
			}
		}
	}
}
