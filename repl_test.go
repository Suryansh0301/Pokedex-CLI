package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "welcome to the pokemon center  ",
			expected: []string{"welcome", "to", "the", "pokemon", "center"},
		},
		{
			input:    "  I am nurse joy  ",
			expected: []string{"I", "am", "nurse", "joy"},
		},
		{
			input:    "we hope to  see you again!",
			expected: []string{"we", "hope", "to", "see", "you", "again!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Error("expected length is not equal to the actual")
			t.Errorf("expected: %v , actual: %v", c.expected, actual)
			t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected word doesn't match the actual word, expected: %s actual: %s", expectedWord, word)
				t.FailNow()
			}
		}
	}
}
