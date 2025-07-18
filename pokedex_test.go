package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	var EMPTY_STRING_SLICE = []string{}
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: EMPTY_STRING_SLICE,
		},
		{
			input: "You WiLl nOt lIVe To sEE the dawn.",
			expected: []string{"you", "will", "not", "live", "to", "see", "the", "dawn."},
		},
		{
			input: "7 \nmiNUTeS \n is \n all \n I can\n spare",
			expected: []string{"7", "minutes", "is", "all", "i", "can", "spare"},
		},
		{
			input: "\n\n\n\n\n\n",
			expected: EMPTY_STRING_SLICE,
		},
		{
			input: "            ",
			expected: EMPTY_STRING_SLICE,
		},
		{
			input: "    \n     \n    \n",
			expected: EMPTY_STRING_SLICE,
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual and expected not equal: %v != %v", len(actual), len(c.expected))
			t.Fail()
		}
		for i, word := range (actual) {
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Received %s, expected %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
