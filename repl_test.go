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
			input:    "Hello how are   you?",
			expected: []string{"hello", "how", "are", "you?"},
		},
		{
			input:    "A i U e O",
			expected: []string{"a", "i", "u", "e", "o"},
		},
		{
			input:    "I AM NOT ANGRY",
			expected: []string{"i", "am", "not", "angry"},
		},
		{
			input:    "Pablo is the master of test writting",
			expected: []string{"pablo", "is", "the", "master", "of", "test", "writting"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			errLen := fmt.Errorf("The amount of elements doesnt match, Actual: %v, Expected: %v", len(actual), len(c.expected))
			fmt.Println(errLen)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				errWord := fmt.Errorf("Expected: %v != Got: %v", expectedWord, word)
				fmt.Println(errWord)
				t.Fail()
			}
		}
	}
}
