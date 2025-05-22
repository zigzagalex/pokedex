package main

import "testing"



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
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual doesn't match expected: len actual %v, len expected %v",len(actual), len(c.expected))
			t.Fail()
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words don't match: actual word %s, expected word %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}

